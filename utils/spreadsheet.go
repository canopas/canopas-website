package utils

/**
Terminologies:
    Spreadsheet: The primary object in Google Sheets that can contain multiple sheets, each with structured information contained in cells.
    Sheet: A page or tab within a spreadsheet.
**/

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"

	b64 "encoding/base64"

	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// 3 months total hours
const QUARTER_HOURS = 2160

func (repo *utilsRepository) SaveJobsToSpreadSheet(records []string) {

	ctx := context.Background()

	sheetSrv := googleServiceAuth(ctx, "https://www.googleapis.com/auth/spreadsheets", true)

	if sheetSrv == nil {
		return
	}

	srv := sheetSrv.(*sheets.Service)

	/**
	spreadsheet url = https://docs.google.com/spreadsheets/d/<SPREADSHEETID>/edit#gid=<SHEETID>
	*/

	spreadsheetId := os.Getenv("JOBS_SPREADSHEET_ID")
	request := []*sheets.Request{}
	currentTime := time.Now()

	newYearStarted, newQuarterStarted, sheetId := getSheetData(spreadsheetId, srv, ctx)

	if newYearStarted {
		// create new spreadsheet on every year
		createNewSpreadSheet(currentTime, records, srv, ctx)
		return
	} else if newQuarterStarted {
		// create new sheet on every quarter (3 months diff)
		request = newSheetRequest(currentTime, records)
	} else {
		request = []*sheets.Request{
			{
				AppendCells: &sheets.AppendCellsRequest{
					Fields:  "*",
					Rows:    populateCells(false, records),
					SheetId: int64(sheetId),
				},
			},
		}
	}

	// create the batch request
	batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: request,
	}

	// execute the request
	res, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, &batchUpdateRequest).Context(ctx).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}
}

func googleServiceAuth(ctx context.Context, scope string, sheet bool) interface{} {
	credBytes, err := b64.StdEncoding.DecodeString(os.Getenv("RECAPTCHA_CONFIG_JSON_BASE64"))
	if err != nil {
		log.Error(err)
		return nil
	}

	config, err := google.JWTConfigFromJSON(credBytes, scope)
	if err != nil {
		log.Error(err)
		return nil
	}

	client := config.Client(ctx)

	var srv interface{}

	if sheet {
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Error(err)
			return nil
		}

	} else {
		srv, err = drive.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Error(err)
			return nil
		}
	}

	return srv
}

func getSheetData(spreadsheetId string, srv *sheets.Service, ctx context.Context) (bool, bool, int64) {
	/** get latest sheet title and sheetID,
	    compare its month with current month
		create new one, if difference is > 3 months
		create new spreadsheet every year
	**/

	resp, err := srv.Spreadsheets.Get(spreadsheetId).Context(ctx).Do()
	if err != nil {
		log.Error(err)
		return false, false, 0
	}

	sheetId := resp.Sheets[len(resp.Sheets)-1].Properties.SheetId

	prevSheetTitle := resp.Sheets[len(resp.Sheets)-1].Properties.Title

	prevTime, err := time.Parse("January 2006", prevSheetTitle)
	if err != nil {
		log.Error(err)
		return false, false, sheetId
	}

	// create new sheet in spreadsheet on each quarter
	newQuarterStarted := time.Now().Sub(prevTime).Hours() > QUARTER_HOURS

	// create new spreadsheet every year
	newYearStarted := newQuarterStarted && strings.Contains(strings.ToLower(prevSheetTitle), "octo")
	if newYearStarted {
		sheetId = 0
	}

	return newYearStarted, newQuarterStarted, sheetId
}

func populateCells(isTitle bool, records []string) []*sheets.RowData {
	cells := []*sheets.CellData{}
	for i := range records {
		data := &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: &(records[i]),
			},
			UserEnteredFormat: &sheets.CellFormat{
				BackgroundColor: &sheets.Color{ // white background
					Alpha: 1,
					Blue:  1,
					Red:   1,
					Green: 1,
				},
				TextFormat: &sheets.TextFormat{
					Bold: isTitle,
				},
			},
		}
		cells = append(cells, data)
	}

	return []*sheets.RowData{
		{Values: cells},
	}
}

func createNewSpreadSheet(currentTime time.Time, records []string, srv *sheets.Service, ctx context.Context) {
	sheetId := int64(rand.Intn(10))

	// create spreadsheet
	sheetResp, err := srv.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: "Jobs Applications " + strconv.Itoa(currentTime.Year()),
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{
					SheetId: sheetId,
					Title:   currentTime.Month().String() + " " + strconv.Itoa(currentTime.Year()),
				},
				Data: []*sheets.GridData{
					{
						StartRow: 0,
						RowData:  populateCells(true, titleRecords()),
					},
					{
						StartRow: 2,
						RowData:  populateCells(false, records),
					},
				},
			},
		},
	}).Context(ctx).Do()

	if err != nil {
		log.Error(err)
		return
	}

	spreadsheetId := sheetResp.SpreadsheetId

	// add permission to the sheet as its created by client_email
	driveSrv := googleServiceAuth(ctx, "https://www.googleapis.com/auth/drive.file", false)

	if driveSrv == nil {
		return
	}

	dSrv := driveSrv.(*drive.Service)

	_, err = dSrv.Permissions.Insert(spreadsheetId, &drive.Permission{
		Role:  "writer",
		Type:  "user",
		Value: os.Getenv("JOBS_RECEIVER"),
	}).Context(ctx).Do()

	if err != nil {
		log.Error(err)
		return
	}

	os.Setenv("JOBS_SPREADSHEET_ID", spreadsheetId)
}

func newSheetRequest(currentTime time.Time, records []string) []*sheets.Request {
	sheetId := int64(rand.Intn(10))

	return []*sheets.Request{
		{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheets.SheetProperties{
					SheetId: sheetId,
					Title:   currentTime.Month().String() + " " + strconv.Itoa(currentTime.Year()),
				},
			},
		},
		{
			AppendCells: &sheets.AppendCellsRequest{
				Fields:  "*",
				Rows:    populateCells(true, titleRecords()),
				SheetId: sheetId,
			},
		},
		{
			AppendCells: &sheets.AppendCellsRequest{
				Fields:  "*",
				Rows:    populateCells(false, records),
				SheetId: sheetId,
			},
		},
	}
}

func titleRecords() []string {
	return []string{"Name", "Email", "Mobile", "Date", "Location", "Applied For", "Experience", "Reference", "Interview", "Status", "Reason for Rejection", "Offered Salary"}
}
