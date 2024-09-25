<script lang="ts">
    import { page } from "$app/stores";
    import type {
        PollStateResponse,
        PollStateDetails,
        AggregateOption,
    } from "$lib/ambient";
    import { pollState } from "$lib/index";
    import { quadInOut } from "svelte/easing";
    import { fade } from "svelte/transition";

    const pollId = $page.params.pollId;
    let loading: boolean = $state(false);
    let loadingText: string = $state("");

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
        loadingText = "Loading result...";
        pollState(pollId)
            .then((obj) => {
                let pollStateResponse: PollStateResponse = obj;
                pollResult = pollStateResponse.result.result;
                question = pollResult.question;
                options = pollResult.options;
                optionKeys = Object.keys(options);
                totalVotes = pollResult.totalVotes;
            })
            .catch((error) => {
                window.alert("Unable to fetch poll details : " + error);
            })
            .finally(() => {
                loading = false;
            });
    });

    function loadHome() {
        location.href = `/`;
    }
</script>

{#if loading || options.length == 0}
    <span
        class="flex flex-col justify-center text-center text-slate-600 text-4xl font-extrabold min-h-screen"
        >{loadingText}</span
    >
{:else}
    <div
        class="flex flex-col justify-center text-left min-h-screen max-w-lg m-auto p-6"
    >
        <div class="text-slate-600 text-6xl mx-auto mb-6 font-extrabold">
            Result
        </div>
        <span class="text-left text-slate-600 text-2xl rounded-lg border-2 border-gray-500 p-4 bg-white drop-shadow-md font-bold w-full mb-8"
            >Que - {question}</span
        >
        <!-- <span class="text-left text-slate-400 text-md w-full mb-8">
    Poll ID : {pollId}
</span> -->
        <div class="flex flex-col w-full mb-8 space-y-4">
            <span class="flex items-left text-slate-500 text-xl font-extrabold">
                Total Votes : {totalVotes}
            </span>
            <div class="flex flex-col gap-2 w-full">
                {#each optionKeys as option}
                    <button
                        in:fade={{
                            duration: 300,
                            easing: quadInOut,
                        }}
                        class="bg-white text-gray-700 border-2 border-indigo-500 w-full p-4 text-center rounded-lg font-bold transition duration-300 ease-in-out hover:drop-shadow-lg hover:text-white hover:bg-indigo-500"
                    >
                        {option} [{options[option]["total"]}/{totalVotes}];
                    </button>
                {/each}
            </div>
        </div>
        <button
            class="rounded-lg bg-indigo-500 text-white font-medium w-full h-12 hover:bg-indigo-600 transition duration-300 ease-in-out"
            onclick={loadHome}
        >
            Go to Home
        </button>
    </div>
{/if}
