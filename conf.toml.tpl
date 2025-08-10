[mysql]
host = "127.0.0.1"
port = 3306
pretable = "mi_"
username = "root"
password = "1Qazxsw2"
database = "cloud"
protocol = "mysql"
charset = "utf8mb4"

[app]
tag = ["json"]
path = "./tmp"
file = true
null_or_sql = "sql"
type_package = "database/sql"  # "database/sql" || "github.com/guregu/null"
def_package = "com"
use_table_name = true
use_table_columns = true
kw_prefix = "kw_"
keyword = ["import","case"]
app_key = "44"
app_id = "1234"
app_secret = "1234"

[http]
scheme = "http" 
root = "dist"
index = "index.html"
port = 20327
domain = "proxy.51d.ink"
server = "console.51d.ink"   # if has port :port,eg: a.com:8080

[ws]
port = 8080
domain = "ws.51d.ink"

# Error in creating struct from json: error formatting: 1:9: expected 'IDENT', found 'case' (and 1 more errors), was formatting