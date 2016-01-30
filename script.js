(function(){
    var apply = function(e) {
        if (e.getAttribute('rel') != 'noreferrer') return;
        var href = e.getAttribute('href');
        if (!href || href.match(/^(?:https?:)?\/\/fw\.mearie\.org\/./)) {
            return;
        } else if (href.match(/^http:/)) {
            href = href.substr(7);
        } else if (href.match(/^https:/)) {
            href = '*' + href.substr(8);
        } else {
            return;
        }
        if (href.match(/^[^\/]+$/)) href += '/';
        e.setAttribute('href', '//fw.mearie.org/' + href);
    };

    if (document.querySelectorAll) {
        var ee = document.querySelectorAll('a[rel=noreferrer]'), l = ee.length;
        for (var i = 0; i < l; ++i) apply(ee[i]);
    } else if (document.evaluate) {
        var ee = document.evaluate('//a[@rel="noreferrer"]', document, null, XPathResult.UNORDERED_NODE_ITERATOR_TYPE, null), e;
        while (e = ee.iterateNext) apply(e);
    } else {
        var ee = document.getElementsByTagName('a'), l = ee.length;
        for (var i = 0; i < l; ++i) {
            if (ee[i].getAttribute('rel') == 'noreferrer') apply(ee[i]);
        }
    }
}());
