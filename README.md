Results for 1M rows
===================

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


PY 2.7.6
--------
    K = 100000
    0.68970418  d1 list
    0.02689290  d2 list

    --- start join
    0.03684616  d2 hash
    0.74472117  d1 part of result hash
    0.04486609  add d2 part to result hash
    --- full join time: 0.826585054398

    1.58251095  format result_hast to table


PY 3.3.5
--------
    K = 100000
    0.58259225  d1 list
    0.02027702  d2 list

    --- start join
    0.03363848  d2 hash
    0.71283388  d1 part of result hash
    0.04388213  add d2 part to result hash
    --- full join time: 0.7905449867248535

    0.47126508  format result_hast to table

