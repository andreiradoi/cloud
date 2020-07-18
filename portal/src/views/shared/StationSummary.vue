<template>
    <div class="station-hover-summary" v-if="viewingSummary && station">
        <div class="upper-half">
            <div class="row general-row">
                <div class="image-container">
                    <img alt="Station Image" :src="stationSmallPhoto" class="" />
                </div>
                <div class="station-details">
                    <div class="station-name">{{ station.name }}</div>
                    <div class="station-synced">
                        Last Synced
                        <span class="small-light">{{ getSyncedDate() }}</span>
                    </div>
                    <div class="station-battery">
                        <img class="battery" alt="Battery Level" :src="getBatteryImg()" />
                        <span class="small-light">{{ station.battery }}</span>
                    </div>
                    <div v-for="module in station.modules" v-bind:key="module.id" class="module-icon-container">
                        <img v-if="!module.internal" alt="Module Icon" class="small-space" :src="getModuleIcon(module)" />
                    </div>
                </div>
                <div class="close-button" v-on:click="wantCloseSummary">
                    <img alt="Close" src="@/assets/close.png" />
                </div>
            </div>

            <div class="row where-row">
                <div class="location-container" v-if="station.locationName ? station.locationName : station.placeNameOther">
                    <img alt="Location" src="@/assets/icon-location.png" class="icon" />
                    <template>{{ station.locationName ? station.locationName : station.placeNameOther }}</template>
                </div>
                <div class="location-container" v-else>
                    <img alt="Location" src="@/assets/icon-location.png" class="icon" />
                    <template>Awaiting Location</template>
                </div>

                <div class="location-container" v-if="station.placeNameNative">
                    <img alt="Location" src="@/assets/icon-location.png" class="icon" />
                    <template>Native Lands: {{ station.placeNameNative }}</template>
                </div>
                <div class="location-container" v-else>
                    <img alt="Location" src="@/assets/icon-location.png" class="icon" />
                    <template>Native Lands: Awaiting Location</template>
                </div>

                <div class="coordinates-row">
                    <div v-if="station.location" class="location-coordinates valid">
                        <div class="coordinate latitude">
                            {{ station.location.latitude | prettyCoordinate }}
                            <br />
                            Latitude
                        </div>
                        <div class="coordinate longitude">
                            {{ station.location.longitude | prettyCoordinate }}
                            <br />
                            Longitude
                        </div>
                    </div>
                    <!--
                    <div v-else class="location-coordinates missing">
                        <div class="coordinate latitude">
                            --
                            <br />
                            Latitude
                        </div>
                        <div class="coordinate longitude">
                            --
                            <br />
                            Longitude
                        </div>
                    </div>
					-->
                    <div class="empty"></div>
                </div>
            </div>
            <div class="readings-container" v-if="readings">
                <div class="title">Latest Readings</div>
                <LatestStationReadings :id="station.id" />
            </div>
            <div class="explore-button" v-if="readings" v-on:click="onClickExplore">
                Explore Data
            </div>
        </div>
    </div>
</template>

<script>
import _ from "lodash";
import * as utils from "@/utilities";
import { makeAuthenticatedApiUrl } from "@/api/api";
import { BookmarkFactory } from "@/views/viz/viz";
import CommonComponents from "@/views/shared";

export default {
    name: "StationSummary",
    components: {
        ...CommonComponents,
    },
    data: () => {
        return {
            viewingSummary: true,
        };
    },
    props: {
        station: {
            type: Object,
            required: true,
        },
        readings: {
            type: Boolean,
            default: true,
        },
    },
    computed: {
        stationSmallPhoto() {
            return makeAuthenticatedApiUrl(this.station.photos.small);
        },
    },
    methods: {
        viewSummary() {
            this.viewingSummary = true;
        },
        onClickExplore() {
            const bm = BookmarkFactory.forStation(this.station.id);
            return this.$router.push({ name: "exploreBookmark", params: { bookmark: JSON.stringify(bm) } });
        },
        getSyncedDate() {
            return utils.getUpdatedDate(this.station);
        },
        getBatteryImg() {
            const battery = this.station.battery;
            let img = "";
            if (battery == 0) {
                img = "0.png";
            } else if (battery <= 20) {
                img = "20.png";
            } else if (battery <= 40) {
                img = "40.png";
            } else if (battery <= 60) {
                img = "60.png";
            } else if (battery <= 80) {
                img = "80.png";
            } else {
                img = "100.png";
            }
            return this.$loadAsset("battery/" + img);
        },
        getSensorName(module, sensor) {
            const newName = utils.convertOldFirmwareResponse(module);
            return this.$t(newName + ".sensors." + sensor.name);
        },
        getModuleIcon(module) {
            return this.$loadAsset(utils.getModuleImg(module));
        },
        wantCloseSummary() {
            this.$emit("close");
        },
    },
};
</script>

<style scoped>
.station-hover-summary {
    position: absolute;
    background-color: #ffffff;
    border: 1px solid rgb(215, 220, 225);
    z-index: 2;
    display: flex;
    flex-direction: column;
    padding: 20px;
}
.image-container {
    width: 124px;
    height: 100px;
    text-align: center;
}
.image-container img {
    max-width: 124px;
    max-height: 100px;
}
.station-name {
    font-size: 18px;
    font-weight: bold;
}
.station-synced {
    font-size: 14px;
}
.battery {
    width: 20px;
    height: 11px;
}
.module-icon-container {
    float: left;
    margin-right: 5px;
}
.location-coordinates {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    flex-basis: 50%;
}
.location-coordinates .coordinate {
    display: flex;
    flex-direction: column;
}
.general-row {
    display: flex;
    flex-direction: row;
}
.coordinates-row {
    display: flex;
    flex-direction: row;
}
.where-row > div {
    margin-top: 10px;
}
.where-row {
    border-top: 1px solid #f1eeee;
    margin-top: 20px;
    padding-top: 10px;
    display: flex;
    flex-direction: column;
}
.close-button {
    margin-left: auto;
    cursor: pointer;
}
.station-details {
    text-align: left;
}
.where-row {
    text-align: left;
    font-size: 14px;
    color: #2c3e50;
}
.readings-container {
    margin-top: 20px;
    padding-top: 10px;
    border-top: 1px solid #f1eeee;
    text-align: left;
    font-size: 14px;
    color: #2c3e50;
}
.readings-container div.title {
    padding-bottom: 10px;
}
.explore-button {
    font-size: 18px;
    font-weight: bold;
    color: #ffffff;
    text-align: center;
    padding: 10px;
    margin: 20px 0 0 0px;
    background-color: #ce596b;
    border: 1px solid rgb(215, 220, 225);
    border-radius: 4px;
    cursor: pointer;
}
</style>