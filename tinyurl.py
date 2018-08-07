import random
class Codec:
    def __init__(self):
        self.map = dict()
        self.baselen = len("http://tinyurl.com/")
    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.
        
        :type longUrl: str
        :rtype: str
        """
        k = random.randint(1, 65535)
        self.map[k] = longUrl
        return "http://tinyurl.com/" + str(k)  

    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.
        
        :type shortUrl: str
        :rtype: str
        """
        return self.map[int(shortUrl[self.baselen:])]
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))
