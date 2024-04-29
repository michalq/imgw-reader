# Skierniewice Draft

# Chart no.1
<iframe title="Average temperature - yearly" aria-label="Interactive line chart" id="datawrapper-chart-2CFEn" src="https://datawrapper.dwcdn.net/2CFEn/2/" scrolling="no" frameborder="0" style="width: 0; min-width: 100% !important; border: none;" height="429" data-external="1"></iframe><script type="text/javascript">!function(){"use strict";window.addEventListener("message",(function(a){if(void 0!==a.data["datawrapper-height"]){var e=document.querySelectorAll("iframe");for(var t in a.data["datawrapper-height"])for(var r=0;r<e.length;r++)if(e[r].contentWindow===a.source){var i=a.data["datawrapper-height"][t]+"px";e[r].style.height=i}}}))}();
</script>
```sqlite
select y as year, avg(t)
from measurements
group by y;
```

# Chart no.2
<iframe title="Days with higher avg temp than 10°C, 20°C, 25°C " aria-label="Interactive area chart" id="datawrapper-chart-LMWkr" src="https://datawrapper.dwcdn.net/LMWkr/1/" scrolling="no" frameborder="0" style="width: 0; min-width: 100% !important; border: none;" height="450" data-external="1"></iframe><script type="text/javascript">!function(){"use strict";window.addEventListener("message",(function(a){if(void 0!==a.data["datawrapper-height"]){var e=document.querySelectorAll("iframe");for(var t in a.data["datawrapper-height"])for(var r=0;r<e.length;r++)if(e[r].contentWindow===a.source){var i=a.data["datawrapper-height"][t]+"px";e[r].style.height=i}}}))}();
</script>

```sqlite
select y as year,
       sum(tmax>10) t10, sum(tmax>20) t20, sum(tmax>25) t25
from measurements
group by y
```

# Chart no.3
Monthly distribution pivot.

<iframe title="Monthly distribution" aria-label="Interactive line chart" id="datawrapper-chart-PYVFN" src="https://datawrapper.dwcdn.net/PYVFN/1/" scrolling="no" frameborder="0" style="width: 0; min-width: 100% !important; border: none;" height="645" data-external="1"></iframe><script type="text/javascript">!function(){"use strict";window.addEventListener("message",(function(a){if(void 0!==a.data["datawrapper-height"]){var e=document.querySelectorAll("iframe");for(var t in a.data["datawrapper-height"])for(var r=0;r<e.length;r++)if(e[r].contentWindow===a.source){var i=a.data["datawrapper-height"][t]+"px";e[r].style.height=i}}}))}();
</script>

```sqlite
with years as (
  select distinct y as year
  from measurements
),
lines as (
  select 'select m ' as part
  union all
  select ', avg(t) filter (where y = ' || year || ') as "' || year || '" '
  from years
  union all
  select 'from measurements group by m order by m;'
)
select group_concat(part, '')
from lines;
```

# Chart no.4
<iframe title="Days in year where min temp &amp;lt; 0" aria-label="Interactive line chart" id="datawrapper-chart-iAo11" src="https://datawrapper.dwcdn.net/iAo11/1/" scrolling="no" frameborder="0" style="width: 0; min-width: 100% !important; border: none;" height="430" data-external="1"></iframe><script type="text/javascript">!function(){"use strict";window.addEventListener("message",(function(a){if(void 0!==a.data["datawrapper-height"]){var e=document.querySelectorAll("iframe");for(var t in a.data["datawrapper-height"])for(var r=0;r<e.length;r++)if(e[r].contentWindow===a.source){var i=a.data["datawrapper-height"][t]+"px";e[r].style.height=i}}}))}();
</script>

```sqlite
select y as year,
       sum(tmin<0) tl0
from measurements
group by y
```