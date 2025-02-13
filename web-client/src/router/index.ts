import {
	createRouter,
	createWebHistory,
	type RouteRecordRaw
} from "vue-router";
import { authGuard } from "@auth0/auth0-vue";
import notFound from "../pages/not-found.vue";
import home from "../pages/home.vue";
import authCallback from "../pages/auth-callback.vue";
import profileGuard from "./guards/profile-guard";
import CreateProfile from "../pages/create-profile.vue";
import UpdateProfile from "../pages/update-profile.vue";

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: "home",
		component: home
	},
	{
		path: "/create-profile",
		name: "create-profile",
		component: CreateProfile,
		beforeEnter: authGuard
	},
	{
		path: "/update-profile",
		name: "update-profile",
		component: UpdateProfile,
		beforeEnter: authGuard
	},
	{
		path: "/callback",
		name: "callback",
		component: authCallback
	},
	{
		path: "/:catchAll(.*)",
		name: "Not Found",
		component: notFound
	}
];

const router = createRouter({
	history: createWebHistory(),
	routes
});

router.beforeEach(profileGuard);

export default router;
