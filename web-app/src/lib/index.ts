import { PUBLIC_API_KEY } from '$env/static/public';

import { v4 as uuid } from 'uuid'

const apiKey = PUBLIC_API_KEY;

const callApi = async (endpoint: string, args: { [key: string]: any }) => {
    const params = {
        network: 'TESTNET',
        blockchain: 'KALP',
        walletAddress: 'b1036fdd6fa97e145dac4b874056aaefdb7da1e3',
        args: args,
    };

    try {
        const response = await fetch(endpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': apiKey!,
            },
            body: JSON.stringify(params),
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.message || 'Something went wrong');
        }
        return data;
    } catch (err: any) {
        throw err;
    }
};

export async function createPoll (id: string, question: string, options: string[], locations: string[], durationHours: number) {
    const endpoint =
        'https://gateway-api.kalp.studio/v1/contract/kalp/invoke/Xm89gsGTpPCp0vPtnSNUmLgQvstScy551727007362011/CreatePoll';
    const args = {
        pollId: id,
        question: question,
        options: options,
        locations: locations,
        durationHours: durationHours
    };
    return callApi(endpoint, args);
};

export async function getPoll (pollId: string) {
    const endpoint =
        'https://gateway-api.kalp.studio/v1/contract/kalp/query/Xm89gsGTpPCp0vPtnSNUmLgQvstScy551727007362011/PollState';
    const args = { pollId: pollId };
    return callApi(endpoint, args);
};

export async function submitVote (pollId: string, optionIndex: number, locationIndex: number) {
    const endpoint =
        'https://gateway-api.kalp.studio/v1/contract/kalp/invoke/Xm89gsGTpPCp0vPtnSNUmLgQvstScy551727007362011/SubmitVote';
    const args = {
        pollId: pollId,
        optionIndex: optionIndex,
        locationIndex: locationIndex
    };
    return callApi(endpoint, args);
};

export async function pollState (pollId: string) {
    const endpoint =
        'https://gateway-api.kalp.studio/v1/contract/kalp/query/Xm89gsGTpPCp0vPtnSNUmLgQvstScy551727007362011/PollState';
    const args = { pollId: pollId };
    return callApi(endpoint, args);
};

export async function isPollOpen (pollId: string) {
    const endpoint =
        'https://gateway-api.kalp.studio/v1/contract/kalp/query/Xm89gsGTpPCp0vPtnSNUmLgQvstScy551727007362011/IsPollOpen';
    const args = {
        pollId: pollId
    };
    return callApi(endpoint, args);
};