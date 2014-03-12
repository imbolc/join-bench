import sys
import time
from pprint import pprint

if sys.version.startswith('3'):
    xrange = range

K = 1000 * 1000
K = 100 * 1000

print('K = %s' % K)

t = time.time()
d1 = [[1, i, 10] for i in xrange(10 * K)]
print('%0.8f\t%s' % (time.time() - t, 'd1 list'))

t = time.time()
d2 = [[1, i, 5] for i in xrange(2 * K, 3 * K)]
print('%0.8f\t%s' % (time.time() - t, 'd2 list'))

print('\n--- start join')
tjoin = time.time()

t = time.time()
d2_hash = {(a, b): c for a, b, c in d2}
print('%0.8f\t%s' % (time.time() - t, 'd2 hash'))

t = time.time()
result_hash = {(a, b): (c, d2_hash.get((a, b))) for a, b, c in d1}
print('%0.8f\t%s' % (time.time() - t, 'd1 part of result hash'))


t = time.time()
for a, b, c in d2:
    if (a, b) not in result_hash:
        result_hash[(a, b)] = (a, b, None, c)
print('%0.8f\t%s' % (time.time() - t, 'add d2 part to result hash'))

print('--- full join time: %s\n' % (time.time() - tjoin))

t = time.time()
result = [(k[0], k[1], v[0], v[1]) for k, v in result_hash.items()]
print('%0.8f\t%s' % (time.time() - t, 'format result_hast to table'))

print('\n')
pprint(result[:10])
