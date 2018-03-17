<template>
  <v-app id="inspire" dark>
    <v-navigation-drawer
    clipped
    fixed
    v-model="drawer"
    app
    >
    <v-list dense>
      <v-list-tile @click="">
        <v-list-tile-action>
          <v-icon>dashboard</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Dashboard</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile @click="">
        <v-list-tile-action>
          <v-icon>settings</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Settings</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
  <v-toolbar app fixed clipped-left>
    <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
    <v-toolbar-title>Application</v-toolbar-title>
  </v-toolbar>
  <v-content>
    <v-container fluid>
      <v-layout justify-center align-center>
        <v-flex xs12>
          <v-select
          placeholder="First or last name"
          autocomplete
          :loading="loading"
          :items="items"
          solo
          :search-input.sync="search"
          v-model="select"
          color="light-green accent-3"
          dense
          @input="OnChange"
          ></v-select>
          <vheatmap
          year=2010
          cellSize=12>
            
          </vheatmap> 
          <!-- required -->
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
  <v-footer app fixed>
    <span>&copy; 2017</span>
  </v-footer>
</v-app>
</template>

<script>
import axios from 'axios';
import vheatmap from './components/vheatmap.vue';
export default {
  data: () => ({
    drawer: true,
    search: null,
    items: [],
    select: [],
    selected: null,
    loading: false
  }),
  watch: {
    search (val) {
      val && this.querySelections(val)
    }
  },
  methods: {
    OnChange(){
      this.items = [];
      axios.get("/api/list/"+ this.select)
      .then((response)  =>  {
        this.loading = false;
          console.log(response.data);
        }, (error)  =>  {
          this.loading = false;
          console.log(error);
        });

      console.log(this.select + "changed");
    },
    querySelections (v) {
      if (!v || v.length < 3 ){
        this.items = [];
        return;
      }
      this.loading = true;
        // Simulated ajax query
        axios.get("/api/search?query="+v)
        .then((response)  =>  {
          this.loading = false;
          this.items = response.data;
          // console.log(response.data);
        }, (error)  =>  {
          this.loading = false;
          console.log(error);
        })

      }
    },
    components: {
      'vheatmap': vheatmap
    },
    props: {
      source: String
    }
  }
  </script>
