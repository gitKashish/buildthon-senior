<script lang="ts">
  import "../app.css";
  let {children} = $props();
</script>

<nav class="bg-indigo-400 text-white">Layout Navbar</nav>
{@render children()}
<footer class= "text-slate-400">Powered by KALP DLT</footer>