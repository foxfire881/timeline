{{- define "public/page.gtpl"}}
<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8"/>
    <title>{{.appname}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge"/>
    <link rel="stylesheet" href="/static/public/amis/sdk/sdk.css"/>
    <link rel="stylesheet" href="/static/public/amis/sdk/helper.css"/>
    <link rel="stylesheet" href="/static/public/amis/sdk/iconfont.css"/>

    <style>
        html,
        body,
        .app-wrapper {
            position: relative;
            width: 100%;
            height: 100%;
            margin: 0;
            padding: 0;
        }
    </style>
</head>
<body>
<div id="root" class="app-wrapper"></div>
<script src="/static/public/amis/sdk/sdk.js"></script>
<script type="text/javascript">
    window.enableAMISDebug = false;
    (function () {
        const xhr = new XMLHttpRequest();
        xhr.open('GET',"/static/{{.page}}.json",false);
        xhr.send()
        let amis = amisRequire('amis/embed');
        let amisScoped = amis.embed('#root', JSON.parse(xhr.responseText),{
            responseAdaptor(api, response, query, request) {
                return response;
            }
        });
    })();
</script>
</body>
</html>
{{end}}