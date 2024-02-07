export interface Promotion {
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
