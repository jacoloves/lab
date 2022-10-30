import imp


import http.server, socketserver
from logging import Handler

Handler = http.server.SimpleHTTPRequestHandler
Handler.extensions_map['.wasm'] = 'application/wasm'

port = 8888
with socketserver.TCPServer(("", port), Handler) as d:
    print("[Start] http://localhost:{}".format(port))
    try:
        d.serve_forever()
    except:
        pass
    finally:
        d.server_close()