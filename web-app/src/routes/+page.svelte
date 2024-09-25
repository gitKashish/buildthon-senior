<script lang="ts">
    import { isPollOpen } from "$lib/index";

    let pollId = $state("");

    let loading = $state(false);

    function loadPollPage() {
        pollId = pollId.trim();
        if (pollId.length != 36) {
            alert("Invalid Poll ID");
            pollId = "";
            loading = false;
            return;
        }
        loading = true;
        isPollOpen(pollId)
            .then((response) => {
                const isOpen = response.result.result;
                if (isOpen) {
                    location.href = "/poll/" + pollId;
                } else {
                    location.href = "/result/" + pollId;
                }
            })
            .catch(() => {
                window.alert("Poll with id " + pollId + "does not exist.");
                pollId = "";
            })
            .finally(() => {
                loading = false;
            });
    }

    function loadCreatePage() {
        location.href = "/create";
    }
</script>

<div class="flex flex-col gap-4 justify-center items-center min-h-screen p-6">
    <div
        class="flex flex-col justify-center items-center text-slate-600 text-6xl min-h-20 font-extrabold"
    >
        DeVote
        <div class="flex justify-center items-center text-slate-500 text-xl">
            Decentralized Voting
        </div>
    </div>
    <div class="w-full max-w-sm space-y-4">
        <input
            id="pollIdField"
            class="rounded-lg border bg-white border-gray-300 p-3 w-full focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:drop-shadow-md transition-shadow ease-in-out duration-200 disabled:bg-gray-200 disabled:text-gray-600"
            type="text"
            placeholder="Enter the Poll ID"
            required
            bind:value={pollId}
            disabled={loading}
        />
        <button
            class="rounded-lg bg-indigo-500 text-white font-medium w-full h-12 hover:bg-indigo-600 transition duration-300 ease-in-out  disabled:bg-indigo-200 disabled:text-indigo-500"
            onclick={loadPollPage}
            disabled={loading}
        >
            {loading ? "Loading..." : "Open Poll"}
        </button>
        <button
            class="rounded-lg border border-indigo-500 text-indigo-500 font-medium w-full h-12 hover:bg-indigo-500 hover:text-white transition duration-300 ease-in-out disabled:bg-indigo-200 disabled:text-indigo-500 disabled:border-0"
            onclick={loadCreatePage}
            disabled={loading}
        >
            {loading ? "Loading..." : "Create Poll"}
        </button>
    </div>
</div>
