import { watch } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { ApiError } from "../../services/api";
import type { NavigationGuard } from "vue-router";
import { getProfile } from "../../services/profiles";
import { profileStore } from "../../stores/profile";

const profileGuard: NavigationGuard = async (to, _from, next) => {
	const { isAuthenticated, user, isLoading } = useAuth0();

	// Wait for isLoading to become false
	await new Promise<void>(resolve => {
		if (!isLoading.value) {
			resolve(); // Resolve immediately if already loaded
		} else {
			const stopWatching = watch(isLoading, newValue => {
				if (!newValue) {
					stopWatching(); // Stop watching once it's false
					resolve();
				}
			});
		}
	});

	// If the user is authenticated, check the profile
	if (isAuthenticated.value && user.value?.sub && !profileStore.profile) {
		try {
			const profile = await getProfile(user.value.sub);
			profileStore.setProfile(profile);
		} catch (error) {
			if (error instanceof ApiError) {
				if (error.statusCode === 404 && to.path !== "/create-profile") {
					return next({ path: "/create-profile" });
				}
			} else {
				console.error(error);
			}
		}
	}

	// Proceed to the requested route
	next();
};

export default profileGuard;
