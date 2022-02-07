> 上面TCP版和HTTP版,数据编码采用的都是默认的gob编码，而gob编码是Go特有的编码和解码的专用序列化方法，这也就意味着Gob无法跨语言使用。而JSON RPC则可以跨语言使用。

> http://liuqh.icu/2022/01/10/go/rpc/02-kuai-su-shi-yong/
