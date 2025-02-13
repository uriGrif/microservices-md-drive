import { reactive } from "vue";
import type { Profile } from "../shared/types";

export const profileStore = reactive({
	profile: null as null | Profile,
	setProfile(profile: Profile) {
		this.profile = profile;
	},
	clearProfile() {
		this.profile = null;
	}
});
