<template>
  <v-container fluid>
    <v-card v-if="CalendarData.ptotal + CalendarData.utotal > 0">
     <v-card-text >
      <!-- <div class="text-xs-center"> -->
        <h2 class="text-xs-center">{{CalendarData.title}}</h2>
        <!-- </div> -->
      </v-card-text>
      <v-card-media>
        <v-layout row>
          <v-flex class="calender-map">
            <svg viewBox="0 0 700 135" xmlns="http://www.w3.org/2000/svg" v-if="CalendarData.year">
              <!-- <text transform="translate(350,10)" font-size="10px"  font="sans-serif" width="100%" style="text-anchor: middle;" fill="black">{{CalendarData.title}}</text> -->
              <g transform="translate(40,30)" 
              font-size="10px" font="sans-serif" shape-rendering="crispEdges" width="100%">
              <text transform="translate(-25,42)rotate(-90)" style="text-anchor: middle;  fill: white;">{{CalendarData.year}}</text>

              <text class="dow" v-for="d in 7" :transform="getDayTranslations(d)"  dy="-.25em"     fill="black" text-anchor="end">{{dayNames[d-1]}}</text>

              <g v-for="i in 12" class="month" :transform="getTranslation(i)" >
                <text style="text-anchor: end;" dy="-.25em">{{monthNames[i-1]}}</text>
              </g>
              <rect  v-for="d in AllDays" 
              stroke="#666"
              text-anchor="end"
              :width="cSize" :height="cSize" 
              :x="getX(d)" :y="getY(d)" :fill="getColor(d)" :fill-opacity="getOpacity(d)">
              <title v-if="isLeaveDay(d)">{{getHoverText(d)}}</title>
            </rect>

            <path v-for="m in 12" :d="getPath(m)"
            class="border" fill="none" stroke="#000" stroke-width="2px"></path>

            <!-- <text x="1.6em" dy="13.3em" fill="black" font-size="7px" font-weight="bold" v-if="showPlanned && CalendarData.pcounts.length > 0">Planned:</text> -->

            <!-- <text x="1.6em" dy="14.8em" fill="black" font-size="7px" font-weight="bold" v-if="showUnplanned && CalendarData.ucounts.length > 0">Unplanned:</text> -->

           <!--  <g v-for="(pl, ind) in CalendarData.pcounts" v-if="showPlanned">
              <rect stroke="#666" :width="cSize/2" :height="cSize/2" 
              :x="getLegendx(ind)+'em'" y="8.7em" :fill="colors[pl.cat_id-1]"> </rect>
              <text :x="getLegendText(ind)+'em'" dy="13.3em" font-size="7px" fill="black">{{pl.cat}}</text>
            </g>
            <g v-for="(ul, ind) in CalendarData.ucounts" v-if="showUnplanned">
              <rect  stroke="#666"  :width="cSize/2" :height="cSize/2" 
              :x="getLegendx(ind)+'em'" y="9.8em" :fill="colors[ul.cat_id-1]"> </rect> 
              <text :x="getLegendText(ind)+'em'" dy="14.8em" fill="black" font-size="7px">{{ul.cat}}</text>
            </g> -->
          </g>
        </svg>
      </v-flex>
    </v-layout>
  </v-card-media>
  <!-- <v-card-actions> -->
    <v-flex xs2 offset-xs10>
      <v-radio-group v-model="viewMode" column>
        <v-radio label="All" value="total" color="blue darken-3"></v-radio>
        <v-radio label="Planned" value="planned" color="green darken-3"></v-radio>
        <v-radio label="Unplanned" value="unplanned" color="red darken-2"></v-radio>
      </v-radio-group>
    </v-flex>
  <!-- </v-card-actions> -->
  </v-card>
</v-flex> 
</v-layout>
</v-card-text>
</v-card>
</v-container>
</template>

<script>
import moment from "moment";
import numeral from "numeral";

export default {
  data: () => ({
    showPlanned: true,
    showUnplanned: true,
    legend: {},
    test: "hello",
    plannedDays:{},
    unplannedDays:{},
    viewMode: "total", // "planned","unplanned","total" 
    AllDays: [],

    cSize :12,
    firstSunday: null,
    dayNames:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
    monthNames : [ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" ],
    colors : { 
      total:'#0D47A1',
      unplanned:'#D32F2F', //'#BA68C8',
      planned:'#1B5E20'}
  }),
  methods: {
    countDays:function(days){
      var dcount = 0;
      for (var i = 0; i< days.length; i++){
        dcount += days[i].count;
      }
      return dcount;
    },
    countAll: function(upls , pls){
      return this.countDays(upls) + this.countDays(pls);
    },
    getLegendx: function(index,planned=true){
      return 10*(index)+5;
    },
    getLegendText: function(index,planned=true){
      return 10*(index)+ index * 4.3 + 8.4;
    },
    weekNo: function(daDate){
      return Math.ceil((((daDate - this.firstSunday) / 86400000) + 1)/7);
    },
    getX: function(daDate){
      return this.cSize * this.weekNo(daDate);
    },
    getY: function(daDate){
      return this.cSize * daDate.getDay();
    },
    isLeaveDay: function(daDate){
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      switch(this.viewMode){
        case "planned":
          return ( kw in this.plannedDays  && this.plannedDays[kw] > 0 );
          break;
        case "unplanned":
          return (  kw in  this.unplannedDays  && this.unplannedDays[kw] >0 );
          break;
        case "total":
        default:
          return ( kw in this.plannedDays  && this.plannedDays[kw] > 0 ) || (  kw in  this.unplannedDays  && this.unplannedDays[kw] >0 );
          break;
      }
      
    },
    getHoverText: function(daDate){
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var count = 0;
      var max = 0;
      var min = 0;
      switch(this.viewMode){
        case "planned":
          count = this.plannedDays[kw];
          min = this.CalendarData.pmin;
          max = this.CalendarData.pmax;
          break;
        case "unplanned":
          count = this.unplannedDays[kw];
          min = this.CalendarData.umin;
          max = this.CalendarData.umax;
          break;
        case "total":
        default:
        if (this.unplannedDays[kw]==undefined && this.plannedDays[kw]==undefined) {
          count = 0;
          min = 0;
          max = 0;
        } else if (this.plannedDays[kw]==undefined){
          count = this.unplannedDays[kw];
          min = this.CalendarData.umin;
          max = this.CalendarData.umax;
        }else if(this.unplannedDays[kw]==undefined){
          count = this.plannedDays[kw];
          min = this.CalendarData.pmin;
          max = this.CalendarData.pmax;
        }else{
          count = this.plannedDays[kw] + this.unplannedDays[kw];
          min = Math.min(this.CalendarData.umin, this.CalendarData.pmin);
          max = this.CalendarData.umax + this.CalendarData.pmax;
          break;
        }
          
      }
      return "Count: " + count + ", Range: [" + min + ".." + max + "]" ;
    },
    getColor: function(daDate) {
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var z = 0.0;
      var color = this.colors[this.viewMode];

      if( this.viewMode === 'planned' &&  ( kw in  this.plannedDays ) && this.plannedDays[kw] > 0 ){
        return color;
      }
      if( this.viewMode === 'unplanned' && ( kw in  this.unplannedDays ) && this.unplannedDays[kw] >0 ){
        return color;
      }
      if ( this.viewMode === 'total' && ( kw in  this.plannedDays || kw in  this.unplannedDays ) ){
        return color;
      }
      if ( m.weekday()%6==0){
          return "#595959";
        }
      return "#424242";
    },
    getOpacity: function(daDate) {
      var m = moment(daDate);
      var kw = m.format("DD-MM");
      var count = 0;
      var min = 0, max = 0;
      switch(this.viewMode){
        case "planned":
          max = this.CalendarData.pmax;
          min = this.CalendarData.pmin;
          if (kw in this.plannedDays){
            count = this.plannedDays[kw];
          }
          break;
        case "unplanned":
          max = this.CalendarData.umax;
          min = this.CalendarData.umin;
          if (kw in this.unplannedDays){
            count = this.unplannedDays[kw];
          }
          break;
        case "total":
        default:
          max = Math.max(this.CalendarData.pmax,this.CalendarData.umax) ;
          min = Math.min(this.CalendarData.pmin,this.CalendarData.umin) ;
          if (kw in this.plannedDays){
            count = this.plannedDays[kw];
          }
          if (kw in this.unplannedDays){
            count += this.unplannedDays[kw];
          }
      } 

      if (count == 0){
        return 1.0;
      }
      var base = 0.22;
      return (count - min)/(max - min) * (1 - base) + base  ;
    },
    getTranslation: function(i){
      return "translate("+ (8+ i* 4.3 * this.cSize) + ",-3)";
    },
    getDayTranslations: function(d){
      return "translate(0,"+ d * this.cSize + ")";
    },
    getPath: function(m) {
      var t0 = new Date(Date.UTC(this.CalendarData.year, m - 1 , 1));
      var t1 = new Date(Date.UTC(this.CalendarData.year, m , 0));
      var d0 = t0.getDay();
      var d1 = t1.getDay();
      var w0 = this.weekNo(t0);
      var w1 = this.weekNo(t1);
      return "M" + (w0 + 1) * this.cSize + "," + d0 * this.cSize
      + "H" + w0 * this.cSize + "V" + 7 * this.cSize
      + "H" + w1 * this.cSize + "V" + (d1 + 1) * this.cSize
      + "H" + (w1 + 1) * this.cSize + "V" + 0
      + "H" + (w0 + 1) * this.cSize + "Z";
    },
    generateDates: function() {
      console.log("===>>",this.CalendarData.year);
      var fromDate = new Date(Date.UTC(this.CalendarData.year, 0, 1));
      var ToDate = new Date(Date.UTC(this.CalendarData.year,11,31));
      this.AllDays.push(fromDate);
      while (fromDate < ToDate) {
        fromDate = new Date(Date.UTC(fromDate.getUTCFullYear(), fromDate.getUTCMonth(), fromDate.getUTCDate() + 1));
        this.AllDays.push(fromDate);
      }
    }
  },
  created: function () {
    this.generateDates();
    var fd_dow = new Date(Date.UTC(this.CalendarData.year,0,1)).getDay();
    this.firstSunday = new Date(Date.UTC(this.CalendarData.year, 0, 1 - fd_dow));
    this.plannedDays = this.CalendarData.pdays;
    this.unplannedDays = this.CalendarData.udays;
    
  },
  props:['CalendarData', 'LeaveTypes']
}
</script>

