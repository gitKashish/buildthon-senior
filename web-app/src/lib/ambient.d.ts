export interface PollResponse {
    timestamp: string;
    status: string;
    message: string;
    result: {
        success: boolean;
        httpStatus: number;
        result: PollDetails;
    };
}

export interface PollDetails {
    question: string;
    options: string[];
    locations: string[];
    expiry: string;
    votes: Votes[]; // Define this more specifically if the structure of votes is known
}

export interface Votes {
    pollId: string;
    optionIndex: number;
    locationIndex: number;
}

export interface PollStateResponse {
    timestamp: string;
    status: string;
    message: string;
    result: {
        success: boolean;
        httpStatus: number;
        result: PollStateDetails;
    };
}

export interface PollStateDetails {
    pollId: string;
    question: string;
    totalVotes: number;
    options: AggregateOption[]
}

export interface AggregateOption {
    [option: string] : LocationFrequency[]
}

export interface LocationFrequency {
    [key: string] : number;
}