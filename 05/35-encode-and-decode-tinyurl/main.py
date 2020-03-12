class Codec:
    def __init__(self):
        self.mapping = {}
        self.back = {}
        self.next = 0

    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.

        :type longUrl: str
        :rtype: str
        """
        if longUrl in self.mapping:
            return "http://tinyurl.com/" + str(self.mapping[longUrl])

        id = self.next
        self.next += 1
        self.mapping[longUrl] = id
        self.back[id] = longUrl
        return "http://tinyurl.com/" + str(self.mapping[longUrl])


    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.

        :type shortUrl: str
        :rtype: str
        """
        id = int(shortUrl[len("http://tinyurl.com/"):])
        return self.back[id]


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))
