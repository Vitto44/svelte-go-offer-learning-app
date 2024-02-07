/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		extend: {
			colors: {
				primary: "#004080",
				secondary: {
					100: "#FFA300",
					200: "#FFA500",
				},
				text: {
					primary: "#ffffff",
					gray: "#666666",
				},
			},
		},
	},
	plugins: [],
};
