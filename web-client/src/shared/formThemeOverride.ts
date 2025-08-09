import type { GlobalThemeOverrides } from "naive-ui";

const themeOverrides: GlobalThemeOverrides = {
	Form: {
		labelTextColor: "var(--text-color)"
	},
	Input: {
		color: "var(--primary-light-color)",
		colorFocusError: "var(--text-color)",
		borderRadius: "var(--border-radius)",
		border: "0px",
		textColor: "var(--primary-dark-color)",
		placeholderColor: "var(--primary-dark-color)",
		colorFocus: "var(--text-color)",
		borderHover: "1px solid var(--accent-color)",
		borderFocus: "1px solid var(--accent-color)",
		caretColor: "var(--secondary-color)"
	},
	Select: {
		peers: {
			InternalSelection: {
				arrowColor: "var(--primary-dark-color)",
				color: "var(--primary-light-color)",
				borderRadius: "var(--border-radius)",
				border: "0px",
				textColor: "var(--primary-dark-color)",
				placeholderColor: "var(--secondary-color)",
				borderHover: "1px solid var(--accent-color)",
				borderFocus: "1px solid var(--accent-color)",
				borderActive: "1px solid var(--accent-color)"
			},
			InternalSelectMenu: {
				color: "var(--primary-light-color)",
				borderRadius: "5px",
				optionTextColor: "var(--primary-dark-color)",
				optionTextColorActive: "var(--primary-dark-color)"
			}
		}
	}
};

export default themeOverrides;
