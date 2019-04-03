<template lang="html">
    <div class="container">
        Hello {{bundler}}
        <button v-on:click="counter = counter + 1">Click</button>
        {{counter}}
    </div>
</template>


<script lang="ts">

    import Vue from "vue";

    export default Vue.extend({
        data() {
            return {
                bundler: "Parcel and Vue",
                counter: 0
            };
        },
        methods: {
            get: function (event) {

            }
        }
    });
</script>

<style lang="scss" scoped>
    .container {
        color: green;
    }
</style>