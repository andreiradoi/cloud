<template>
    <div class="project-container" v-if="visible">
        <router-link :to="{ name: 'viewProject', params: { id: project.id } }">
            <div class="project-image-container">
                <!--
                <div v-if="invited" class="invited-icon">
                    Project Invite
                </div>
				-->
                <ProjectPhoto :project="project" />
            </div>

            <img v-if="project.private" alt="Private project" src="@/assets/private.png" class="private-icon" />

            <div class="project-name">{{ project.name }}</div>
            <div class="project-description">{{ project.description }}</div>
            <div class="invited-container" v-if="invited">
                <div class="accept" v-on:click.stop.prevent="onAccept">
                    <img alt="Close" src="@/assets/tick.png" />
                    Accept Invite
                </div>
                <div class="reject" v-on:click.stop.prevent="onDecline">
                    <img alt="Close" src="@/assets/Delete.png" />
                    Decline
                </div>
            </div>
            <div class="social-container" v-else>
                <div class="social follows" v-if="project.following">
                    <img alt="Follows" src="@/assets/heart.png" class="follow-icon" />
                    <span>{{ project.following.total }}</span>
                </div>
                <!--
                <div class="social notifications" v-if="project.notifications">
                    <img alt="Notifications" src="@/assets/notification.png" class="notify-icon" />
                    <span>2</span>
                </div>
                <div class="social comments" v-if="project.comments">
                    <img alt="Comments" src="@/assets/comment.png" class="comment-icon" />
                    <span>3</span>
                </div>
				-->
            </div>
        </router-link>
    </div>
</template>
<script lang="ts">
import CommonComponents from "@/views/shared";
import * as ActionTypes from "@/store/actions";

export default {
    name: "ProjectThumbnail",
    components: {
        ...CommonComponents,
    },
    props: {
        project: {
            type: Object,
            required: true,
        },
        invited: {
            type: Boolean,
            default: false,
        },
    },
    data() {
        return {
            visible: true,
            accepted: false,
        };
    },
    methods: {
        getImageUrl(this: any, project) {
            return this.$config.baseUrl + project.photo;
        },
        onAccept(this: any, ev: any) {
            console.log("accept", ev);
            return this.$store.dispatch(ActionTypes.ACCEPT_PROJECT, { projectId: this.project.id }).then(() => {
                this.accepted = true;
                this.visible = false;
            });
        },
        onDecline(this: any, ev: any) {
            console.log("decline", ev);
            return this.$store.dispatch(ActionTypes.DECLINE_PROJECT, { projectId: this.project.id }).then(() => {
                this.visible = false;
            });
        },
    },
};
</script>

<style scoped>
.project-container {
    width: 270px;
    height: 265px;
    border: 1px solid #d8dce0;
    margin-right: 20px;
    margin-bottom: 20px;
}
.project-name {
    font-weight: bold;
    font-size: 16px;
    margin: 10px 15px 0 15px;
}
.project-description {
    overflow-wrap: break-word;
    font-weight: lighter;
    font-size: 14px;
    margin: 0 15px 10px 15px;
}
.project-image-container {
    height: 138px;
    text-align: center;
    border-bottom: 1px solid #d8dce0;
}
/deep/ .project-image {
    max-width: 270px;
    max-height: 138px;
}
.invited-icon {
    float: right;
    position: relative;
    z-index: 10;
    top: 0;
    right: 0;
}
.private-icon {
    float: right;
    margin: -14px 14px 0 0;
    position: relative;
    z-index: 10;
}
.social-container {
}
.social {
    display: inline-block;
    font-size: 14px;
    font-weight: 600;
    margin: 0 14px 0 15px;
}
.social img {
    float: left;
    margin: 2px 4px 0 0;
}
.invited-container {
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    padding: 10px;
}
.invited-container .accept {
    flex: 1;
    padding: 5px;
    text-align: center;
    font-size: 14px;
    font-weight: 900;
    color: #2c3e50;
    border-radius: 3px;
    border: solid 1px #cccdcf;
}
.invited-container .reject {
    flex: 1;
    padding: 5px;
    text-align: center;
    font-size: 14px;
    font-weight: 900;
    color: #2c3e50;
}
</style>