<template>
  <div v-if="build" class="flex text-gray-600 dark:text-gray-500">
    <BuildStatusIcon :build="build" class="flex items-center" />
    <div class="flex flex-col ml-4">
      <span class="underline">{{ build.owner }} / {{ build.name }}</span>
      <span>{{ message }}</span>
      <div class="flex flex-col mt-2">
        <div class="flex space-x-2 items-center">
          <Icon name="duration" />
          <span>{{ duration }}</span>
        </div>
        <div class="flex space-x-2 items-center">
          <Icon name="since" />
          <span>{{ since }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, toRef } from 'vue';

import Icon from '~/components/atomic/Icon.vue';
import BuildStatusIcon from '~/components/repo/build/BuildStatusIcon.vue';
import useBuild from '~/compositions/useBuild';
import { BuildFeed } from '~/lib/api/types';

export default defineComponent({
  name: 'BuildFeedItem',

  components: { BuildStatusIcon, Icon },

  props: {
    build: {
      type: Object as PropType<BuildFeed>,
      required: true,
    },
  },

  setup(props) {
    const build = toRef(props, 'build');
    const { since, duration, message } = useBuild(build);

    return { since, duration, message };
  },
});
</script>
