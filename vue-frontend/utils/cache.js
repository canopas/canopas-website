class Cache {
  constructor(config) {
    config = config || {};
    config.ttl = config.ttl || 1 * 3600;

    var data = {};
    var self = this;

    var now = function () {
      return new Date().getTime() / 1000;
    };

    /**
     * Object for holding a value and an expiration time
     * @param expires the expiry time as a UNIX timestamp
     * @param value the value of the cache entry
     * @constructor ¯\(°_o)/¯
     */
    var CacheEntry = function (expires, value) {
      this.expires = expires;
      this.value = value;
    };

    /**
     * Creates a new cache entry with the current time + ttl as the expiry.
     * @param value the value to set in the entry
     * @returns {CacheEntry} the cache entry object
     */
    CacheEntry.create = function (value) {
      return new CacheEntry(now() + config.ttl, value);
    };

    /**
     * Clears all cache entries.
     */
    this.clear = function () {
      for (var key in data)
        if (Object.prototype.hasOwnProperty.call(data, key)) self.remove(key);
    };

    /**
     * Gets the cache entry for the given key.
     * @param key the cache key
     * @returns {*} the cache entry if set, or undefined otherwise
     */
    this.get = function (key) {
      if (!Object.prototype.hasOwnProperty.call(data, key)) {
        return null;
      }

      let entry = data[key];

      if (this.expired(entry.expires, now())) {
        self.remove(key);
        return null;
      }

      return entry.value;
    };

    /**
     * Sets a cache entry with the provided key and value.
     * @param key the key to set
     * @param value the value to set
     */
    this.put = function (key, value) {
      data[key] = CacheEntry.create(value);
    };

    /**
     * Removes the cache entry for the given key.
     * @param key the key to remove
     */
    this.remove = function (key) {
      delete data[key];
    };

    /**
     * Checks if the cache entry has expired.
     * @param entrytime the cache entry expiry time
     * @param curr (optional) the current time
     * @returns {boolean} true if expired, false otherwise
     */
    this.expired = function (entrytime, curr) {
      if (!curr) curr = now();
      return entrytime < curr;
    };
  }
}

export default Cache;
