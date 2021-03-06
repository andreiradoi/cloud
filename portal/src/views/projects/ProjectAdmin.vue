<template>
    <div class="project-admin project-container" v-if="project">
        <div class="details">
            <div class="left">
                <div class="photo-container">
                    <ProjectPhoto :project="project" />
                </div>

                <DisplayProjectTags :tags="project.tags" />
                <FollowPanel :project="project" v-bind:key="project.id" />
            </div>

            <div class="right">
                <div class="details-heading">
                    <div class="title">Project Details</div>
                    <div v-on:click="editProject" class="link">Edit Project</div>
                </div>
                <div class="details-top">
                    <div class="details-left">
                        <div class="project-detail" v-if="project.goal">Project Goal: {{ project.goal }}</div>
                        <div class="project-detail">{{ project.description }}</div>
                    </div>
                    <div class="details-right">
                        <div class="details-row" v-if="project.startTime">
                            <div class="details-icon-container">
                                <img alt="Location" src="@/assets/icon-calendar.svg" class="icon" width="14px" height="14px" />
                            </div>
                            <template>Started: {{ project.startTime | prettyDate }}</template>
                        </div>
                        <div class="details-row" v-if="displayProject.duration">
                            <div class="details-icon-container">
                                <img alt="Location" src="@/assets/icon-time.svg" class="icon" width="14px" height="14px" />
                            </div>
                            <template>{{ displayProject.duration | prettyDuration }}</template>
                        </div>
                        <div class="details-row" v-if="project.location" width="12px" height="14px">
                            <div class="details-icon-container">
                                <img alt="Location" src="@/assets/icon-location.svg" class="icon" />
                            </div>
                            <template>{{ project.location }}</template>
                        </div>
                        <div class="details-row" v-if="displayProject.places.native">
                            <div class="details-icon-container">
                                <img alt="Location" src="@/assets/icon-location.svg" class="icon"  width="12px" height="14px" />
                            </div>
                            <template>Native Lands: {{ displayProject.places.native }}</template>
                        </div>
                    </div>
                </div>

                <div class="details-bottom">
                    <div class="details-team">
                        <div class="title">Team</div>
                        <template v-for="projectUser in displayProject.users">
                            <UserPhoto :user="projectUser.user" v-if="!projectUser.invited" v-bind:key="projectUser.user.email" />
                        </template>
                    </div>
                    <div class="details-modules">
                        <div class="title">Modules</div>

                        <img
                            v-for="module in projectModules"
                            v-bind:key="module.name"
                            alt="Module icon"
                            class="module-icon"
                            :src="module.url"
                        />
                    </div>
                </div>
            </div>
        </div>

        <div class="row-section project-stations">
            <ProjectStations :project="project" :admin="true" :userStations="userStations" />
        </div>

        <div class="row-section data-readings">
            <div class="project-data">
                <ProjectDataFiles :projectStations="displayProject.stations" />
            </div>
            <div class="project-readings">
                <StationsReadings :project="project" :displayProject="displayProject" />
            </div>
        </div>

        <TeamManager :displayProject="displayProject" v-bind:key="displayProject.id" />

    </div>
</template>

<script>
import CommonComponents from "@/views/shared";
import ProjectStations from "./ProjectStations.vue";
import ProjectActivity from "./ProjectActivity.vue";
import ProjectDataFiles from "./ProjectDataFiles.vue";
import StationsReadings from "./StationsReadings.vue";
import TeamManager from "./TeamManager.vue";

import * as utils from "../../utilities";

export default {
    name: "ProjectAdmin",
    components: {
        ...CommonComponents,
        ProjectStations,
        ProjectDataFiles,
        StationsReadings,
        TeamManager,
    },
    props: {
        displayProject: {
            type: Object,
            required: true,
        },
        userStations: {
            type: Array,
            required: true,
        },
    },
    data: () => {
        return {
            viewingActivityFeed: false,
        };
    },
    computed: {
        project() {
            return this.displayProject.project;
        },
        projectModules() {
            return this.displayProject.modules.map((m) => {
                return {
                    name: m.name,
                    url: this.getModuleImg(m),
                };
            });
        },
    },
    methods: {
        getProjectUserImage(projectUser) {
            if (projectUser.user.photo) {
                return this.$config.baseUrl + "/" + projectUser.user.photo.url;
            }
            return null;
        },
        openProjectNotes() {
            return this.$router.push({ name: "viewProjectNotes", params: { projectId: this.project.id } });
        },
        editProject() {
            return this.$router.push({ name: "editProject", params: { id: this.project.id } });
        },
        addUpdate() {
            return this.$router.push({ name: "addProjectUpdate", params: { project: this.project } });
        },
        viewProfile() {
            return this.$emit("viewProfile");
        },
        closeActivityFeed() {
            this.viewingActivityFeed = false;
        },
        openActivityFeed() {
            this.viewingActivityFeed = true;
        },
        getModuleImg(module) {
            return this.$loadAsset(utils.getModuleImg(module));
        },
    },
};
</script>

<style scoped lang="scss">
@import '../../scss/mixins';
@import '../../scss/project';

.project-admin {
    display: flex;
    flex-direction: column;
    padding-bottom: 100px;
}
.header {
    display: flex;
    flex-direction: row;
}
.header .left {
    margin-right: auto;
    display: flex;
    flex-direction: column;
}
.header .right {
    margin-left: auto;
}
.header .project-name {
    font-size: 24px;
    font-family: $font-family-bold;
    margin: 0 15px 0 0;
    display: inline-block;
}
.header .project-dashboard {
    font-size: 20px;
    font-family: $font-family-bold;
    margin: 0 15px 0 0;
    display: inline-block;
    margin-top: 10px;
    margin-bottom: 20px;
}

.details {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
}
.details > .left {
    max-width: 400px;
    border: 1px solid #d8dce0;
    border-radius: 1px;
    margin-right: 20px;
    background-color: #fff;
    padding: 25px;
    flex-direction: column;
    justify-content: space-evenly;
    @include flex();

    @include bp-down($md) {
        max-width: unset;
    }

    @include bp-down($xs) {
        flex: 1;
        margin: 0;
        padding: 15px 10px;
    }
}
.details > .photo {
    display: flex;
    flex-direction: column;
    margin-bottom: 10px;
}

.photo-container {
    width: 288px;
    max-height: 139px;
    align-self: center;

    @include bp-down($sm) {
        max-height: 150px;
    }

    @include bp-down($xs) {
        max-height: 150px;
    }
}

.details > .right {
    border: 1px solid #d8dce0;
    border-radius: 1px;
    background-color: white;
    padding: 20px 30px;

    @include bp-down($md) {
        flex-basis: 100%;
        margin-top: 25px;
    }

    @include bp-down($sm) {
        flex-basis: 100%;
        padding: 20px 20px;
    }

    @include bp-down($xs) {
        padding: 14px 10px 20px;
    }
}
.project-stations {
}

.row-section {
}

.details .details-heading {
    display: flex;
    flex-direction: row;
}
.details .details-heading .title {
    padding-bottom: 30px;
    font-size: 20px;
    font-weight: 500;

    @include bp-down($sm) {
        padding-boottom: 25px;
    }
}
.details .details-heading .link {
    margin-left: auto;
}
.link {
    font-family: $font-family-bold;
    font-size: 14px;
    cursor: pointer;
}

.details .details-bottom {
    border-top: 1px solid #d8dce0;
    padding-top: 20px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;

    @include bp-down($xs) {
        padding-top: 15px;
    }
}
.details-bottom .details-team {
    flex: 1;

    @include bp-down($xs) {
        flex-basis: 100%;
        margin-bottom: 15px;
    }
}
.details-bottom .details-modules {
    flex: 1;
}
.details-bottom .title {
    font-weight: 500;
    font-size: 14px;
}
.details-icon-container {
    width: 20px;
    display: flex;
}

.row-section.data-readings {
    margin-top: 25px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
}
.project-data,
.project-readings {
    border: 1px solid #d8dce0;
    border-radius: 1px;
    background-color: white;
    padding: 20px 25px;
    display: flex;
    flex-direction: column;

    @include bp-down($xs) {
        padding: 15px 10px 2px;
    }
}

.project-data {
    margin-right: 20px;
    flex: 2;
    padding: 20px 25px;

    @include bp-down($sm) {
        flex-basis: 100%;
        margin: 0 0 25px;
    }
}
.data-readings .project-readings {
    flex: 1;
    min-width: 360px;

    @include bp-down($xs) {
        min-width: unset;
    }
}

.project-container {
    margin-top: 18px;
}
::v-deep .project-image {
    width: 100%;
    height: auto;
}
.module-icon {
    width: 35px;
    height: 35px;
    margin: 6px 7px 0 0;
}
.project-detail {
    font-family: $font-family-light;

    &:not(:last-of-type) {
        padding-bottom: 6px;
    }
}

::v-deep .default-user-icon {
    margin: 6px 7px 0 0;
    width: 35px;
    height: 35px;
}

</style>
