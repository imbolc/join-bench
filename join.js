var i, k, t, d1, d2, d2_hash, result_hash, result;

k = 1000 * 1000;
k = 100 * 1000;

console.log('K =', k);

t = new Date();
var d1 = [];
for (i = 0; i < 10 * k; i++) {
    d1.push([1, i, 10]);
}
console.log('%s\t%s', (new Date() - t) / 1000, 'd1 list');

t = new Date();
d2 = [];
for (i = 2 * k; i < 3 * k; i++) {
    d2.push([1, i, 5]);
}
console.log('%s\t%s', (new Date() - t) / 1000, 'd2 list');

console.log('\n--- start join');
tjoin = new Date();

t = new Date();
d2_hash = {};
d2.forEach(function (item) {
    d2_hash[item[0] + '-' + item[1]] = item[2];
});
console.log('%s\t%s', (new Date() - t) / 1000, 'd2 hash');

t = new Date();
result_hash = {};
d1.forEach(function (item) {
    var k = item[0] + '-' + item[1];
    result_hash[k] = [item[2], d2_hash[k]];
});
console.log('%s\t%s', (new Date() - t) / 1000, 'd1 part of result hash');

t = new Date();
d2.forEach(function (item) {
    var k = item[0] + '-' + item[1];
    if (!(k in result_hash)) {
        result_hash[k] = [undefined, item[2]];
    }
});
console.log('%s\t%s', (new Date() - t) / 1000, 'add d2 part of result hash');

console.log('--- full join time:', (new Date() - tjoin) / 1000, '\n');

t = new Date();
result = [];
for (k in result_hash) {
    var v = result_hash[k];
    result.push([k, v[0], v[1]]);
}
console.log('%s\t%s', (new Date() - t) / 1000, 'format result_hast to table');

for (i = 0; i < 10 && i < result_hash.length; i++) {
    console.log(result[i]);
}
