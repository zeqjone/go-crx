var body = document.querySelector('body');

init();

function init() {
    var jqTag = createTag('script', '//astyle-src.alicdn.com/app/searchweb/products/zhaohuoshenqi/lib/jquery.js');

    if (jqTag) {
        jqTag.onload = function() {
            loadPluginJs('//astyle-src.alicdn.com/app/searchweb/products/zhaohuoshenqi/entry/js/entry.js');
        };
        body.appendChild(jqTag);
    }
}

function loadPluginJs(src) {
    var pluginTag = createTag('script', src);

    if (pluginTag) {
        body.appendChild(pluginTag);
    }
}

function createTag(type, src) {
    if (type !== 'link' && type !== 'script') {
        return;
    }
    var tag = document.createElement(type);
    if (type === 'link') {
        tag.rel = 'stylesheet';
        tag.href = src;
    } else if (type === 'script') {
        tag.type = 'text/javascript';
        tag.src = src;
    }
    return tag;
}