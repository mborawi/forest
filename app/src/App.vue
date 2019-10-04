<template>
  <v-app>
    <v-app-bar app>
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

    <v-content>
      <v-container fluid>
        <v-row class="justify-space-around">
          <v-col cols="12" sm="6">
            <v-combobox
              v-model="selectedBsl"
              :items="BSLs"
              item-text="name"
              item-value="id"
              label="Business Lines"
              chips
              outlined
              @change="getBranches()"
            ></v-combobox>
          </v-col>
          <v-col cols="12" sm="6">
            <v-combobox
              v-model="selectedCost"
              :items="CostCentres"
              item-text="name"
              label="Cost Centre"
              multiple
              outlined
              :disabled="selectedBsl==null"
              chips
            ></v-combobox>
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
      this.getDepartments();
      this.getLeaves();
  },
  methods:{
    getLeaves(){
      this.leaves = null;
      axios.get("/api/availability/4")
      .then((response)  =>  {
          this.leaves = response.data;
          // this.selectedBsl = null;
        }, (error)  =>  {
          console.log(error);
        });
    },
    getDepartments(){
        this.BSLs = [];
        axios.get("/api/branch/list")
          .then((response)  =>  {
              this.BSLs = response.data;
              // this.selectedBsl = null;
            }, (error)  =>  {
              console.log(error);
            });
    },
    getBranches(){
        this.selectedCost = null;
        this.CostCentres = [];
        axios.get("/api/department/"+this.selectedBsl.id)
          .then((response)  =>  {
              this.CostCentres = response.data;
              // this.selectedBsl = null;
            }, (error)  =>  {
              console.log(error);
            });
    }
  },
  data: () => ({
    //
    leaves: null,
    selectedBsl: null,
    selectedCost: null,
    BSLs : ['All','ATOP','ATOF','EST','SDP'],
    CostCentres: ['u5733', 'u7543', 't6744'],
  }),
};
</script>
