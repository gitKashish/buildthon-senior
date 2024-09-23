<script lang="ts">
    import { page } from "$app/stores";
    const pollId = $page.params.pollId;

    let question: string = "What is your favorite programming language?";
    let options: string[] = [
        "Python",
        "Rust",
        "Go",
        "Haskell mentioned!"
    ];
    let locations: string[] = ["USA", "Europe", "Asia", "Other"];

    let selectedOption: string = "";
    let selectedLocation: string = "";

    function selectOption(option: string) {
        selectedOption = option;
    }

    function selectLocation(location: string) {
        selectedLocation = location;
    }

    async function submitVote() {
        if (!selectedOption || !selectedLocation) {
            window.alert("Please select both an option and a location.");
            return;
        }
        try {
            console.log(
                `Voted for: ${selectedOption} from ${selectedLocation}`,
            );
            window.alert(
                `Your vote for "${selectedOption}" from "${selectedLocation}" has been submitted!`,
            );
        } catch (error) {
            window.alert("Error submitting your vote.");
        }
    }
</script>

<div
    class="flex flex-col justify-center items-center min-h-screen max-w-lg m-auto p-6"
>
    <div class="text-slate-600 text-4xl min-h-20 mb-4 font-extrabold w-full">
        {question}
    </div>
    <span class="text-left text-slate-400 text-lg font-bold w-full mb-8">
        Poll ID : {pollId}
    </span>
    <div class="flex flex-col gap-6 w-full space-y-4">
        <div class="flex flex-col gap-4">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                1. Choose an Option
            </div>
            <div class="flex flex-col gap-2">
                {#each options as option}
                    <button
                        class="{selectedOption === option
                            ? 'bg-indigo-500 drop-shadow-lg text-white'
                            : 'bg-indigo-200'} w-full p-4 text-center rounded-lg font-bold transition border-2 border-transparent duration-300 ease-in-out hover:border-2 hover:border-indigo-500"
                        onclick={() => selectOption(option)}
                    >
                        {option}
                    </button>
                {/each}
            </div>
        </div>

        <div class="flex flex-col gap-4">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                2. Choose Your Location
            </div>
            <div class="flex flex-col gap-2">
                {#each locations as location, i}
                    <button
                        class="{selectedLocation === location
                            ? 'bg-indigo-500 drop-shadow-lg text-white'
                            : 'bg-indigo-200'} w-full p-4 text-center rounded-lg font-bold transition border-2 border-transparent duration-300 ease-in-out hover:border-2 hover:border-indigo-500"
                        onclick={() => selectLocation(location)}
                    >
                        {location}
                    </button>
                {/each}
            </div>
        </div>

        <button
            class="rounded-lg border border-indigo-500 text-indigo-500 font-medium w-full h-12 hover:bg-indigo-500 hover:text-white transition duration-300 ease-in-out"
            onclick={submitVote}
        >
            Submit Vote
        </button>
    </div>
</div>
