// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	interface StoreData {
		location: string;
		type: string;
		niche: string;
	}
	interface Promotion {
		id: number;
		title: string;
		image: string;
		url: string;
		rating: number;
		category: string;
		investment: number;
		reward: number;
		rewardType: string;
		code: string;
		disclaimer: string;
		isMoney?: boolean;
	}
	interface Review {
		id: number;
		title: string;
		image: string;
		url: string;
		rating: number;
		category: string;
		games: string[];
		licensedBy: string;
		availableCountries: string[];
		disclaimer: string;
		welcomeOffer: Promotion;
		description: string;
		howToRegister: string;
	}
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
