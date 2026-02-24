import { watch } from "vue";
import { ApiError } from "../../shared/api";
import type { NavigationGuard } from "vue-router";
import { getProfile } from "../../services/profiles";
import { profileStore } from "../../stores/profile";
import auth0Client from "../../shared/auth0-client";

const profileGuard: NavigationGuard = async (to, _from, next) => {
	const { isAuthenticated, user, isLoading } = auth0Client;

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
