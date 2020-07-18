<template>
    <div class="stations-container">
        <StationPickerModal :stations="userStations" @add="onAddStation" @close="onCloseStationPicker" v-if="addingStation" />
        <div class="section-heading stations-heading">
            FieldKit Stations
            <div class="add-station" v-on:click="showStationPicker" v-if="admin">
                <img src="@/assets/add.png" class="add-station-btn" />
                Add Station
            </div>
        </div>
        <div class="section-body">
            <div class="stations-panel">
                <div v-if="projectStations.length == 0" class="project-stations-no-stations">
                    <h3>No Stations</h3>
                    <p>
                        Add a station to this project to include its recent data and activities.
                    </p>
                </div>
                <div v-if="projectStations.length > 0" class="stations">
                    <TinyStation
                        v-for="station in visibleStations"
                        v-bind:key="station.id"
                        :station="station"
                        :narrow="true"
                        @selected="showSummary(station)"
                    >
                        <div class="station-links">
                            <div class="notes" v-on:click="openNotes(station)">Notes</div>
                            <div class="remove" v-on:click="removeStation(station)">Delete</div>
                        </div>
                    </TinyStation>
                </div>
                <div class="pagination">
                    <PaginationControls :page="page" :totalPages="totalPages" @new-page="onNewPage" />
                </div>
            </div>
            <div class="toggle-icon-container" v-on:click="toggleStations" v-if="false">
                <img v-if="showStationsPanel" alt="Collapse List" src="@/assets/tab-collapse.png" class="toggle-icon" />
                <img v-if="!showStationsPanel" alt="Expand List" src="@/assets/tab-expand.png" class="toggle-icon" />
            </div>
            <div class="project-stations-map-container">
                <StationsMap @mapReady="onMapReady" @showSummary="showSummary" ref="stationsMap" :mapped="mappedProject" />
                <StationSummary v-if="activeStation" :station="activeStation" :readings="false" @close="onCloseSummary" />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import _ from "lodash";
import Vue from "vue";
import * as utils from "../../utilities";
import * as ActionTypes from "@/store/actions";
import FKApi from "@/api/api";
import StationSummary from "@/views/shared/StationSummary.vue";
import StationsMap from "@/views/shared/StationsMap.vue";
import StationPickerModal from "@/views/shared/StationPickerModal.vue";
import TinyStation from "@/views/shared/TinyStation.vue";

import PaginationControls from "@/views/shared/PaginationControls.vue";

export default Vue.extend({
    name: "ProjectStations",
    components: {
        StationSummary,
        StationPickerModal,
        StationsMap,
        PaginationControls,
        TinyStation,
    },
    data: () => {
        return {
            activeStationId: null,
            following: false,
            showStationsPanel: true,
            addingStation: false,
            page: 0,
            pageSize: 4,
        };
    },
    props: {
        project: { required: true },
        admin: { required: true },
        mapContainerSize: { required: true },
        listSize: { required: true },
        userStations: { required: true },
    },
    computed: {
        projectStations(this: any) {
            return this.$getters.projectsById[this.project.id].stations;
        },
        mappedProject(this: any) {
            return this.$getters.projectsById[this.project.id].mapped;
        },
        activeStation(this: any) {
            if (this.activeStationId) {
                return this.$getters.stationsById[this.activeStationId];
            }
            return null;
        },
        visibleStations(this: any) {
            if (!this.projectStations) {
                return [];
            }
            const start = this.page * this.pageSize;
            const end = start + this.pageSize;
            return this.projectStations.slice(start, end);
        },
        totalPages() {
            if (this.projectStations) {
                return Math.ceil(this.projectStations.length / this.pageSize);
            }
            return 0;
        },
    },
    methods: {
        onMapReady(this: any, map) {
            console.log("map ready resize");
            this.map = map;
            this.map.resize();
        },
        showStation(this: any, station) {
            this.$router.push({ name: "viewStation", params: { id: station.id } });
        },
        showStationPicker(this: any) {
            this.addingStation = true;
        },
        onAddStation(this: any, stationId) {
            this.addingStation = false;
            const payload = {
                projectId: this.project.id,
                stationId: stationId,
            };
            return this.$store.dispatch(ActionTypes.STATION_PROJECT_ADD, payload);
        },
        onCloseStationPicker() {
            this.addingStation = false;
        },
        showSummary(station) {
            console.log("showSummay", station);
            this.activeStationId = station.id;
        },
        removeStation(this: any, station) {
            console.log("remove", station);
            if (!window.confirm("Are you sure you want to remove this station?")) {
                return;
            }
            const payload = {
                projectId: this.project.id,
                stationId: station.id,
            };
            return this.$store.dispatch(ActionTypes.STATION_PROJECT_REMOVE, payload);
        },
        openNotes(this: any, station) {
            return this.$router.push({
                name: "viewProjectStationNotes",
                params: {
                    projectId: this.project.id,
                    stationId: station.id,
                },
            });
        },
        onCloseSummary() {
            this.activeStationId = null;
        },
        toggleStations() {
            //
        },
        onNewPage(this: any, page: number) {
            this.page = page;
        },
    },
});
</script>

<style scoped>
.toggle-icon-container {
    float: right;
    margin: 16px -34px 0 0;
    position: relative;
    z-index: 2;
    cursor: pointer;
}
.section {
    float: left;
}
.section-heading {
    font-size: 20px;
    font-weight: 600;
    padding-top: 15px;
    padding-bottom: 15px;
    padding-left: 35px;
    border-bottom: 1px solid lightgray;
}
.stations-heading {
    display: flex;
    flex-direction: row;
}
.section-body {
    display: flex;
    flex-direction: row;
    height: 342px;
}
.stations-container {
    display: flex;
    flex-direction: column;
}
.station-name {
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
}
.add-station {
    margin-left: auto;
    font-size: 14px;
    margin-right: 10px;
    cursor: pointer;
}
.add-station-btn {
    width: 18px;
    vertical-align: text-top;
}
.last-seen {
    font-size: 12px;
    font-weight: 600;
    color: #6a6d71;
}
.stations-container {
    margin: 22px 0 0 0;
    border: 2px solid #d8dce0;
    background-color: #ffffff;
}
.stations-panel {
    transition: width 0.5s;
    flex: 1;
    display: flex;
    flex-direction: column;
}
.stations-panel .stations {
    padding: 10px;
}
.stations-panel .stations > * {
    margin-bottom: 10px;
}
.project-stations-map-container {
    transition: width 0.5s;
    position: relative;
    flex: 2;
}
.station-box {
    height: 38px;
    margin: 20px auto;
    padding: 10px;
    border: 1px solid #d8dce0;
    transition: opacity 0.25s;
}
.project-stations-no-stations {
    width: 80%;
    text-align: center;
    margin: auto;
    padding-top: 20px;
}
.project-stations-no-stations h1 {
}
.project-stations-no-stations p {
}

.station-links {
    margin-left: auto;
    width: 5em;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    border-left: 1px solid #d8dce0;
    padding-left: 10px;
    text-align: center;
}

.station-links .remove {
    cursor: pointer;
    font-size: 12px;
}

.station-links .notes {
    cursor: pointer;
    font-size: 14px;
}

.pagination {
    margin-top: auto;
    padding: 5px;
}

/deep/ .station-hover-summary {
    width: 359px;
    top: 30px;
    left: 122px;
}
</style>