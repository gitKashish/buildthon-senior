<script lang="ts">
import { page } from '$app/stores';

const pollId: string = $page.params.pollId;
</script>

<h1>
    Welcome to result page for Poll {pollId}.
</h1>