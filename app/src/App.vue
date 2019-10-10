<template>
  <v-app>
    <v-app-bar max-height=70>
      <v-toolbar-title class="headline text-uppercase">
        <span>Availability</span>
        <span class="font-weight-light">TOOL</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        text
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
      >
        <span class="mr-2">Gallery</span>
      </v-btn>
    </v-app-bar>

    <v-content >
      <v-container fluid>
        <v-row class="justify-space-around">
          <v-col cols="12" sm="6" class="pb-0">
            <v-combobox
              v-model="selectedBsl"
              :items="BSLs"
              item-text="name"
              item-value="id"
              label="Business Lines"
              outlined
              @change="getCostCentres()"
            ></v-combobox>
          </v-col>
          <v-col cols="12" sm="6" class="pb-0">
            <v-combobox
              v-model="selectedCost"
              :items="CostCentres"
              item-text="name"
              label="Cost Centre"
              outlined
              :disabled="selectedBsl==null"
              @change="CostCentresChanged()"
            ></v-combobox>
          </v-col>
        </v-row>
        <v-row>
              <v-col cols="12" class="pb-0">
                <v-slider
                  color="red"
                  track-color="grey"
                  dense
                  v-model="years"
                  label="Years"
                  min=1
                  max=10
                  ticks
                  thumb-label
                  @change="getLeaves()"
                ></v-slider>
              </v-col>
            </v-row>
      </v-container>
      <TeamHeatMap  v-if="leaves" :CalendarData="leaves" />
    </v-content>
  </v-app>
</template>

<script>
import TeamHeatMap from './components/teamHeatMap';
import axios from 'axios';

export default {
  name: 'App',
  components: {
    TeamHeatMap,
  },
  created:function() {
      this.console = window.console;
      // this.$vuetify.theme.dark = true;
      this.getBSLs();
      this.getLeaves();
  },
  methods:{
    getLeaves(){
      var q ={ years: this.years, bsl: this.selectedBsl.id, cost:this.selectedCost.id};
          // this.leaves = null;
      axios.get("/api/availability", {params:q})
      .then((response)  =>  {
          this.leaves = null;
          this.leaves = response.data;
          // this.selectedBsl = null;
        }, (error)  =>  {
          console.log(error);
        });
    },
    getBSLs(){
        this.BSLs = [{id:0,name:'All'}];
        axios.get("/api/branch/list")
          .then((response)  =>  {
              this.BSLs=response.data;
              this.BSLs.unshift({id:0,name:'All'});
              // this.selectedBsl = null;
            }, (error)  =>  {
              console.log(error);
            });
    },
    CostCentresChanged(){
        console.log("cost Centre changed");
        // this.leaves = null;
        this.getLeaves();

    },
    getCostCentres(){
        this.selectedCost = {id:0,name:'All'};
        this.CostCentres = [{id:0,name:'All'}];
        axios.get("/api/department/"+this.selectedBsl.id)
          .then((response)  =>  {
              this.CostCentres = response.data;
              this.CostCentres.unshift({id:0,name:'All'});
              this.getLeaves();
              // this.selectedBsl = null;
            }, (error)  =>  {
              console.log(error);
            });
    }
  },
  data: () => ({
    //
    years: 5,
    leaves: null,
    selectedBsl: "All",
    selectedCost: 'All',
    BSLs : [],
    CostCentres: ['All'],
  }),
};
</script>
