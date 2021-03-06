<template>
    <StandardLayout>
        <div class="container-wrap notes-view">
            <DoubleHeader :title="project.name" subtitle="Field Notes" backTitle="Back to Dashboard" backRoute="projects" v-if="project" />
            <DoubleHeader title="My Stations" subtitle="Field Notes" backTitle="Back to Dashboard" backRoute="projects" v-if="!project" />

            <div class="lower">
                <div class="loading-container empty" v-if="!hasStations">There are no stations to view.</div>
                <template v-else>
                    <template v-if="isMobileView()">
                        <div class="station-tabs">
                            <div
                                class="tab"
                                v-for="station in stations"
                                v-bind:key="station.id"
                                v-bind:class="{ active: selectedStation.id === station.id }"
                                v-on:click="onSelected(station)"
                            >
                                <div class="tab-wrap">
                                    <div class="name">{{ station.name }}</div>
                                    <div v-if="station.deployedAt" class="deployed">Deployed</div>
                                    <div v-else class="undeployed">Not Deployed</div>
                                </div>
                                <div class="tab-content" v-if="selectedStation && selectedNotes">
                                    <div v-if="loading" class="loading-container">
                                        <Spinner />
                                    </div>
                                    <div class="notifications">
                                        <div v-if="failed" class="notification failed">Oops, there was a problem.</div>

                                        <div v-if="success" class="notification success">Saved.</div>
                                    </div>
                                    <NotesViewer
                                        :station="selectedStation"
                                        :notes="selectedNotes"
                                        v-bind:key="stationId"
                                        v-if="project.project.readOnly"
                                    />
                                    <NotesForm
                                        v-else
                                        :station="selectedStation"
                                        :notes="selectedNotes"
                                        @save="saveForm"
                                        v-bind:key="stationId"
                                        @change="onChange"
                                    />
                                </div>
                                <div v-else class="tab-content empty">Please choose a station from the left.</div>
                            </div>
                        </div>
                    </template>
                    <template v-else>
                        <div class="station-tabs">
                            <div
                                class="tab"
                                v-for="station in stations"
                                v-bind:key="station.id"
                                v-bind:class="{ active: selectedStation.id === station.id }"
                                v-on:click="onSelected(station)"
                            >
                                <div class="tab-wrap">
                                    <div class="name">{{ station.name }}</div>
                                    <div v-if="station.deployedAt" class="deployed">Deployed</div>
                                    <div v-else class="undeployed">Not Deployed</div>
                                </div>
                            </div>
                        </div>
                        <div class="tab-content" v-if="selectedStation && selectedNotes">
                            <div v-if="loading" class="loading-container">
                                <Spinner />
                            </div>
                            <div class="notifications">
                                <div v-if="failed" class="notification failed">Oops, there was a problem.</div>

                                <div v-if="success" class="notification success">Saved.</div>
                            </div>
                            <NotesViewer
                                :station="selectedStation"
                                :notes="selectedNotes"
                                v-bind:key="stationId"
                                v-if="project.project.readOnly"
                            />
                            <NotesForm
                                v-else
                                :station="selectedStation"
                                :notes="selectedNotes"
                                @save="saveForm"
                                v-bind:key="stationId"
                                @change="onChange"
                            />
                        </div>
                        <div v-else class="tab-content empty">Please choose a station from the left.</div>
                    </template>
                </template>
            </div>
        </div>
    </StandardLayout>
</template>

<script lang="ts">
import _ from "lodash";
import Vue from "vue";
import Promise from "bluebird";
import CommonComponents from "@/views/shared";
import StandardLayout from "../StandardLayout.vue";
import StationTabs from "./StationTabs.vue";
import NotesForm from "./NotesForm.vue";
import NotesViewer from "./NotesViewer.vue";

import { mapState, mapGetters } from "vuex";
import * as ActionTypes from "@/store/actions";
import { GlobalState } from "@/store/modules/global";

import { serializePromiseChain } from "@/utilities";

import { Notes, mergeNotes } from "./model";

export default Vue.extend({
    name: "NotesView",
    components: {
        ...CommonComponents,
        StandardLayout,
        NotesForm,
        NotesViewer,
    },
    props: {
        projectId: {
            type: Number,
            required: false,
        },
        stationId: {
            type: Number,
            required: false,
        },
        selected: {
            type: Object,
            required: false,
        },
    },
    data: () => {
        return {
            notes: {},
            dirty: false,
            loading: false,
            success: false,
            failed: false,
            mobileView: window.screen.availWidth < 1040,
        };
    },
    computed: {
        ...mapGetters({ isAuthenticated: "isAuthenticated", isBusy: "isBusy" }),
        ...mapState({
            user: (s: GlobalState) => s.user.user,
            stations: (s: GlobalState) => s.stations.user.stations,
            userProjects: (s: GlobalState) => s.stations.user.projects,
        }),
        hasStations(this: any) {
            return this.visibleStations.length > 0;
        },
        project() {
            if (this.projectId) {
                return this.$getters.projectsById[this.projectId];
            }
            return null;
        },
        visibleStations() {
            if (this.projectId) {
                const project = this.$getters.projectsById[this.projectId];
                if (project) {
                    return project.stations;
                }
                return [];
            }
            return this.$store.state.stations.user.stations;
        },
        selectedStation() {
            if (this.stationId) {
                const station = this.$getters.stationsById[this.stationId];
                if (station) {
                    return station;
                }
            }
            return null;
        },
        selectedNotes() {
            if (this.stationId && this.notes) {
                return this.notes[this.stationId];
            }
            return null;
        },
    },
    watch: {
        stationId(this: any) {
            return this.loadNotes(this.stationId);
        },
    },
    mounted(this: any) {
        const desktopBreakpoint = 768;
        const windowAny: any = window;
        const resizeObserver = new windowAny.ResizeObserver((entries) => {
            console.log("data", this.$data.mobileView);
            const windowWidth = entries[0].contentRect.width;

            if (this.$data.mobileView && windowWidth > desktopBreakpoint) {
                this.$data.mobileView = false;
            }
            if (!this.$data.mobileView && windowWidth < desktopBreakpoint) {
                this.$data.mobileView = true;
            }
        });
        resizeObserver.observe(document.querySelector("body"));

        const pending = [];
        if (this.projectId) {
            pending.push(this.$store.dispatch(ActionTypes.NEED_PROJECT, { id: this.projectId }));
        }
        if (this.stationId) {
            pending.push(this.loadNotes(this.stationId));
        }
        return Promise.all(pending);
    },
    beforeRouteUpdate(this: any, to, from, next) {
        console.log("router: update");
        if (this.maybeConfirmLeave()) {
            next();
        }
    },
    beforeRouteLeave(this: any, to, from, next) {
        console.log("router: leave");
        if (this.maybeConfirmLeave()) {
            next();
        }
    },
    methods: {
        loadNotes(this: any, stationId: number) {
            this.success = false;
            this.failed = false;
            Vue.set(this, "loading", true);
            return this.$services.api.getStationNotes(stationId).then((notes) => {
                Vue.set(this.notes, stationId, notes);
                Vue.set(this, "loading", false);
            });
        },
        onSelected(station) {
            if (this.stationId != station.id) {
                return this.$router.push({
                    name: this.projectId ? "viewProjectStationNotes" : "viewStationNotes",
                    params: {
                        projectId: this.projectId.toString(),
                        stationId: station.id,
                    },
                });
            }
        },
        onChange() {
            this.dirty = true;
        },
        maybeConfirmLeave() {
            if (this.dirty) {
                if (window.confirm("You may have unsaved changes, are you sure you'd like to leave?")) {
                    this.dirty = false;
                    return true;
                } else {
                    return false;
                }
            }
            return true;
        },
        saveForm(this: any, formNotes: Notes) {
            this.success = false;
            this.failed = false;

            return serializePromiseChain(formNotes.addedPhotos, (photo) => {
                return this.$services.api.uploadStationMedia(this.stationId, photo.key, photo.file).then((media) => {
                    console.log(media);
                    return [];
                });
            }).then(() => {
                const payload = mergeNotes(this.notes[this.stationId], formNotes);
                return this.$services.api.patchStationNotes(this.stationId, payload).then(
                    (updated) => {
                        this.dirty = false;
                        this.success = true;
                        console.log("success", updated);
                    },
                    () => {
                        this.failed = true;
                        console.log("failed");
                    }
                );
            });
        },
        isMobileView() {
            return this.$data.mobileView;
        },
    },
});
</script>

<style scoped lang="scss">
@import "../../scss/layout";

.notes-view {
    @include bp-down($md) {
        max-width: 600px;
    }
}
.notes-view .lower {
    display: flex;
    background: white;
    margin-top: 20px;
    position: relative;

    @include bp-down($xs) {
        margin-top: -15px;
    }
}
.loading-container {
    height: 100%;
    @include flex(center);
}
.notes-view .lower .loading-container.empty {
    padding: 20px;
}

.notification.success {
    margin-top: 20px;
    margin-bottom: 20px;
    padding: 20px;
    border: 2px;
    border-radius: 4px;
}
.notification.success {
    background-color: #d4edda;
}
.notification.failed {
    background-color: #f8d7da;
}
.notifications {
    padding: 0 10px;
}

.spinner {
    margin-top: 40px;
    margin-left: auto;
    margin-right: auto;
}

.station-tabs {
    text-align: left;
    display: flex;
    flex-direction: column;
    flex-basis: 250px;
    flex-shrink: 0;
    border-top: 1px solid #d8dce0;
    border-left: 1px solid #d8dce0;
    border-bottom: 1px solid #d8dce0;

    @include bp-down($md) {
        flex-basis: 100%;
    }
}
.tab {
    border-bottom: 1px solid #d8dce0;
    cursor: pointer;

    @include bp-down($md) {
        border-right: 1px solid #d8dce0;
        border-bottom: 0;
    }

    &.active {
        border-left: 4px solid #1b80c9;

        @include bp-down($md) {
            border-left: 0;
        }
    }

    &-wrap {
        position: relative;
        padding: 16px 13px;
        z-index: 10;

        @include bp-down($md) {
            padding: 16px 10px;
            border-right: 0;
            transition: max-height 0.33s;

            &:after {
                background: url("../../assets/icon-chevron-right.svg") no-repeat center center;
                transform: rotate(90deg) translateX(-50%);
                content: "";
                width: 20px;
                height: 20px;
                transition: all 0.33s;
                @include position(absolute, 50% 20px null null);

                .tab.active & {
                    transform: rotate(270deg) translateX(50%);
                }
            }
        }

        .tab.active &:before {
            @include bp-up($md) {
                content: "";
                width: 3px;
                height: 100%;
                background: #fff;
                z-index: $z-index-top;
                @include position(absolute, 0 -2px null null);
            }
        }
    }

    &-content {
        width: calc(100% - 250px);
        z-index: $z-index-top;
        border: 1px solid #d8dce0;

        @include bp-down($md) {
            padding-top: 1px;
            width: 100%;
            max-height: 0;
            border: 0;
            border-top: 1px solid #d8dce0;
            overflow: hidden;

            @at-root .tab.active & {
                max-height: 1000px;
            }
        }
    }
}

.vertical {
    margin-top: auto;
    border-right: 1px solid #d8dce0;
    height: 100%;
}
.name {
    font-size: 16px;
    font-weight: 500;
    color: #2c3e50;
    margin-bottom: 1px;
}
.undeployed {
    @include bp-down($md) {
        padding: 0 10px 0 14px;
        width: calc(100% + 24px);
        box-sizing: border-box;
        margin-left: -14px;
    }
}
.undeployed,
.deployed {
    font-size: 13px;
    color: #6a6d71;
    font-weight: 500;
}

::v-deep textarea {
    overflow-y: hidden;
    resize: none;
    color: #2c3e50;
    font-size: 14px !important;

    @include bp-down($xs) {
        font-size: 12px !important;
    }
}
</style>
