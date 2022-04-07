var body = document.querySelector('body');
var	winFloat = document.getElementById('zhsq-winfloat');
if (winFloat) {
	body.removeChild(winFloat);
} else {
	init();
}

function init() {
	loadJquery();
	loadPluginCss('//astyle-src.alicdn.com/app/searchweb/products/zhaohuoshenqi/winfloat/css/winfloat.css');
}

function loadJquery() {
	var jqTag = createTag('script', '//astyle-src.alicdn.com/app/searchweb/products/zhaohuoshenqi/lib/jquery.js');

	if (jqTag) {
		jqTag.onload = function() {
			loadPluginJs('//astyle-src.alicdn.com/app/searchweb/products/zhaohuoshenqi/winfloat/js/winfloat.js');
		};
		body.appendChild(jqTag);
	}
}

function loadPluginCss(src) {
	var cssTag = createTag('link', src);

	if (cssTag) {
		body.appendChild(cssTag);
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