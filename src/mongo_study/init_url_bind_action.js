//执行脚本 mongo ./init_url_bind_action.js

var url = "mongodb://localhost:27017/iam";
var db = connect(url);
var colName = "url_bind_action"

ret = db.getCollection(colName).createIndex({"url":1},{"unique":true});
if(ret.ok != 1)
{
    printjson(ret)
}

col = db.getCollection(colName)

ret = col.insert([

    //kodo部分
    {
        "url":"kodo:GET:",
        "action":"kodo/ReadObject",
        "desc":"下载文件",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:GET:/buckets",
        "action":"kodo/ListBuckets",
        "desc":"列举一个账号的所有空间",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/mkbucketv2",
        "action":"kodo/CreateBucket",
        "desc":"创建存储空间，同时绑定一个七牛二级域名，用于访问资源",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:GET:/v6/domain/list",
        "action":"kodo/ReadBucket",
        "desc":"获取一个空间绑定的域名列表",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/drop",
        "action":"kodo/ReadBucket",
        "desc":"删除指定存储空间",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/upload",
        "action":"kodo/CreateObject",
        "desc":"用于在一次 HTTP 会话中上传单一的一个文件",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {

        "url":"kodo:POST:/mkblk",
        "action":"kodo/CreateObject",
        "desc":"为后续分片上传创建一个新的块，同时上传第一片数据",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/bput",
        "action":"kodo/CreateObject",
        "desc":"上传指定块的一片数据，具体数据量可根据现场环境调整，同一块的每片数据必须串行上传",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {

        "url":"kodo:POST:/mkfile",
        "action":"kodo/CreateObject",
        "desc":"将上传好的所有数据块按指定顺序合并成一个资源文件",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:GET:/stat",
        "action":"kodo/ReadObject",
        "desc":"仅获取资源的 Metadata 信息，不返回资源内容",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/chgm",
        "action":"kodo/ModifyObject",
        "desc":"修改文件的 MIME 类型信息",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/move",
        "action":"kodo/ReadObject and DeleteObject for src, CreateObject for dest",
        "desc":"将源空间的指定资源移动到目标空间，或在同一空间内对资源重命名",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/copy",
        "action":"kodo/ReadObject for src, CreateObject for dest",
        "desc":"将指定资源复制为新命名资源",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/delete",
        "action":"kodo/DeleteBucket for bucket, DeleteObject for object",
        "desc":"删除指定资源",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:GET:/list",
        "action":"kodo/ListObjects",
        "desc":"用于列举指定空间里的所有文件条目	",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/fetch",
        "action":"kodo/CreateObject",
        "desc":"从指定 URL 抓取资源，并将该资源存储到指定空间中。每次只抓取一个文件，抓取时可以指定保存空间名和最终资源名",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/batch",
        "action":"kodo/Various",
        "desc":"指在单一请求中执行多次获取元信息、移动、复制、删除操作，极大提高资源管理效率",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {

        "url":"kodo:POST:/prefetch",
        "action":"kodo/CreateObject",
        "desc":"对于设置了镜像存储的空间，从镜像源站抓取指定名称的资源并存储到该空间中",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {

        "url":"kodo:POST:/image",
        "action":"kodo/CreateObject",
        "desc":"为存储空间指定一个镜像回源网址，用于取回资源",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {

        "url":"kodo:POST:/deleteAfterDays",
        "action":"kodo/ModifyObject",
        "desc":"设置指定资源的生命周期，即设置一个文件多少天后删除",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {

        "url":"kodo:POST:/chtype",
        "action":"kodo/ModifyObject",
        "desc":"修改文件的存储类型信息，即低频存储和标准存储的互相转换",
        "service":"kodo",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    //fusion部分
    {
        "url":"cdn:GET:/v1/downloadverifyfile/:domain",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:GET:/v1/verifystate",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:GET:/v1/domainstate/:domain",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v1/verify",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/isptraffic",
        "action":"cdn/GetISPReqCount",
        "desc":"各运营商请求次数",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/reqcount",
        "action":"cdn/GetReqCount",
        "desc":"查询区域请求次数",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/regiontraffic",
        "action":"cdn/GetReqCount",
        "desc":"查询区域请求次数",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/topcountip",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/topcounturl",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/trafficip",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/toptrafficip",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/trafficurl",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/toptrafficurl",
        "action":"cdn/GetTop",
        "desc":"TOP查询",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/statuscode",
        "action":"cdn/GetStateCode",
        "desc":"查询状态码",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/hitmiss",
        "action":"cdn/GetHitRate",
        "desc":"查询命中率",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/uniquevisitor",
        "action":"cdn/GetUV",
        "desc":"查询独立IP",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/pageview",
        "action":"cdn/GetUV",
        "desc":"查询独立IP",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/trafficdatatest/v1/traffic/domains/flow",
        "action":"cdn/GetBandwidthAndFlux",
        "desc":"查询流量带宽",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/trafficdatatest/v1/traffic/domains/bandwidth",
        "action":"cdn/GetBandwidthAndFlux",
        "desc":"查询流量带宽",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/trafficdatatest/v1/admin/traffic/user",
        "action":"cdn/Dashboard",
        "desc":"数据总览",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/trafficdatatest/v1/admin/traffic/user/compare",
        "action":"cdn/Dashboard",
        "desc":"数据总览",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/analysedata/v1/portal/downspeed",
        "action":"cdn/GetSpeed",
        "desc":"查询下载速度",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/refresh",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/refresh/list",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/refresh/user/setting",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/refresh/user/setting",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/refresh/user/surplus",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:POST:/prefetch",
        "action":"cdn/Prefetch",
        "desc":"文件预取",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/prefetch/list",
        "action":"cdn/Prefetch",
        "desc":"文件预取",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v2/tune/refresh",
        "action":"cdn/Refresh",
        "desc":"刷新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v2/tune/prefetch",
        "action":"cdn/Prefetch",
        "desc":"文件预取",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v2/tune/bandwidth",
        "action":"cdn/GetBandwidthAndFlux",
        "desc":"查询流量带宽",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v2/tune/flux",
        "action":"cdn/GetBandwidthAndFlux",
        "desc":"查询流量带宽",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v1/deftone/rule",
        "action":"cdn/UpdateThresholdAlarm",
        "desc":"流量带宽阈值告警",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v1/deftone/rule/update",
        "action":"cdn/UpdateThresholdAlarm",
        "desc":"流量带宽阈值告警",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/v1/deftone/rule/:domain",
        "action":"cdn/UpdateThresholdAlarm",
        "desc":"流量带宽阈值告警",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:DELETE:/v1/deftone/rule/:domain",
        "action":"cdn/UpdateThresholdAlarm",
        "desc":"流量带宽阈值告警",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:POST:/domain/:domain",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/pandomain/:domain",
        "action":"cdn/CreateDomain",
        "desc":"创建域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:DELETE:/domain/:domain",
        "action":"cdn/DeleteDomain",
        "desc":"删除域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/domain/:domain",
        "action":"cdn/GetDomainInfo",
        "desc":"查询域名信息",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/domain",
        "action":"cdn/GetDomainList",
        "desc":"获取命名列表",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/domainsearch",
        "action":"cdn/GetDomainList",
        "desc":"获取命名列表",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:GET:/domain/:domain/pandomains",
        "action":"cdn/GetDomainList",
        "desc":"获取域名列表",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/domain/:domain/online",
        "action":"cdn/OnlineDomain",
        "desc":"启用域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:POST:/domain/:domain/offline",
        "action":"cdn/OfflineDomain",
        "desc":"停用域名",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/source",
        "action":"cdn/UpdateSource",
        "desc":"回源配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    {
        "url":"cdn:PUT:/domain/:domain/cache",
        "action":"cdn/UpdateCache",
        "desc":"更新缓存",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/referer",
        "action":"cdn/UpdateReferer",
        "desc":"Referer 防盗链配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/timeacl",
        "action":"cdn/UpdateTimeACL",
        "desc":"时间戳防盗链配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/bsauth",
        "action":"cdn/UpdateBSAuth",
        "desc":"回源鉴权配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/sslize",
        "action":"cdn/UpdateHttps",
        "desc":"HTTPS 配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/httpsconf",
        "action":"cdn/UpdateHttps",
        "desc":"HTTPS 配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/ipacl",
        "action":"cdn/UpdateIpACL",
        "desc":"IP 黑白名单",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:PUT:/domain/:domain/external",
        "action":"cdn/UpdateExternal",
        "desc":"修改特殊配置",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"cdn:POST:/v2/tune/log/list",
        "action":"cdn/DownloadCDNLog",
        "desc":"日志下载",
        "service":"cdn",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    //certificate

    {
        "url":"certificate:GET:/sslcert/:certid",
        "action":"certificate/GetCertificate",
        "desc":"获取证书",
        "service":"certificate",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"certificate:GET:/sslcert",
        "action":"certificate/GetCertificateList",
        "desc":"获取证书列表",
        "service":"certificate",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },

    {
        "url":"certificate:GET:/v3/domains/:domain/search/certs",
        "action":"certificate/GetCertificateList",
        "desc":"获取证书列表",
        "service":"certificate",
        "create_at":new Date('2011-01-01T11:00:00'),
        "update_at":new Date('2011-01-01T11:00:00'),
    },


    ]
);

printjson(ret)







