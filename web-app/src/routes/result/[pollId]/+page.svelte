<script lang="ts">
    import { page } from "$app/stores";
    import type {
        PollStateResponse,
        PollStateDetails,
        AggregateOption,
        LocationFrequency,
    } from "$lib/ambient";
    import { pollState } from "$lib/index";
    import { quadInOut } from "svelte/easing";
    import { fly } from "svelte/transition";

    const pollId = $page.params.pollId;
    let loading: boolean = $state(false);

    let question: string = $state("");
    let options: AggregateOption[] = $state([]);
    let optionKeys: string[] = $state([]);
    let totalVotes: number = $state(0);
    let pollResult: PollStateDetails;

    $effect(() => {
        if (pollId.length != 36) {
            alert("Invalid Poll ID. Returning to home.");
            location.href = `/`;
        }
        loading = true;
        pollState(pollId)
            .then((obj) => {
                let pollStateResponse: PollStateResponse = obj;
                pollResult = pollStateResponse.result.result;
                question = pollResult.question;
                options = pollResult.options;
                optionKeys = Object.keys(options);
                totalVotes = pollResult.totalVotes;
                console.log(options);
            })
            .catch((error) => {
                window.alert("Unable to fetch poll details : " + error);
            })
            .finally(() => {
                loading = false;
            });
    });
</script>

{#if loading}
    <span
        class="flex flex-col justify-center text-center text-slate-600 text-6xl font-extrabold min-h-screen"
        >Loading...</span
    >
{:else}
    <div
        class="flex flex-col justify-center items-center min-h-screen max-w-lg m-auto p-6"
    >
        <div class="text-slate-600 text-4xl font-extrabold w-full">
            Result : {question}
        </div>
        <span class="text-left text-slate-400 text-lg font-bold w-full mb-8">
            Poll ID : {pollId}
        </span>
        <div class="flex flex-col gap-4 w-full items-center">
            <div class="flex text-slate-500 text-xl font-extrabold">
                Total Votes : {totalVotes}
            </div>
            <div class="flex flex-col gap-2 w-full">
                {#each optionKeys as option, i}
                    <button
                        class="bg-white text-gray-700 border-2 border-indigo-500 w-full p-4 text-center rounded-lg font-bold transition duration-300 ease-in-out hover:bg-indigo-500 hover:drop-shadow-lg hover:text-white"
                    >
                        {option} : {options[option]["total"]}
                    </button>
                {/each}
            </div>
        </div>
    </div>
{/if}
