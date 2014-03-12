Results for 1M rows
===================
    py3     : 0.77
    py2     : 0.81
    pypy    : 1.28
    go      : 1.75
    node.js : 2.47

GO
---
    K =  100000
    0.10347967   d1 list
    0.00619498   d2 list

    --- start join
    0.06278591   d2 hash
    1.62941095   d1 part to result hash
    0.06086887   add d2 part to result hash
    --- full join time: 1.75318185

    0.11262346   format result_hast to table


NodeJs v0.10.26
---------------
    K = 100000
    0.279   d1 list
    0.068   d2 list

    --- start join
    0.146   d2 hash
    2.005   d1 part of result hash
    0.311   add d2 part of result hash
    --- full join time: 2.473 

    0.831   format result_hast to table



Python v2.7.6
-------------
    K = 100000
    0.68836093  d1 list
    0.02568698  d2 list

    --- start join
    0.03631616  d2 hash
    0.72987103  d1 part of result hash
    0.04458809  add d2 part to result hash
    --- full join time: 0.810923099518

    1.56242704  format result_hast to table


Python v3.3.5
-------------
    K = 100000
    0.57706189d1 list
    0.02021503d2 list

    --- start join
    0.03339791d2 hash
    0.69553232d1 part of result hash
    0.04203367add d2 part to result hash
    --- full join time: 0.7711665630340576

    0.47044921format result_hast to table


PyPy v2.2.1
-----------
    K = 100000
    0.29356599  d1 list
    0.02470613  d2 list

    --- start join
    0.25577497  d2 hash
    0.99361682  d1 part of result hash
    0.03890204  add d2 part to result hash
    --- full join time: 1.28859591484

    0.53422403  format result_hast to table

