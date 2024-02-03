<!-- component -->
<template>

<article class="border w-2/4 mx-auto border-gray-400 rounded-lg md:p-4 bg-white sm:py-3 py-4 px-2 m-10" data-article-path="#" data-content-user-id="112962">
    <div role="presentation">
        <div>
          <div class="m-2">
            <div class="flex items-center">
                <div class="mr-8">
                </div>
                <div>
                  <p>From: {{ from }}</p>
                  <p>To: {{ to }}</p>
                  <p class="text text-gray-700 hover:text-black">
                    <time datetime="2019-08-02T13:58:42.196Z">Date: {{ date }}</time>
                  </p>
                </div>
            </div>
          </div>
          <div class="pl-12 md:pl-10 xs:pl-10">
            <h2 class="text-2xl mb-2 hover:text-black-500 leading-7">
                <p class="font-bold">
                    Subject: {{ subject }}
                </p>
            </h2>
            <div class="mb-1 leading-6 message-container">
              <div v-if="showLess">
                <p>
                    {{ body !== undefined ? body.substring(0, 130) : '' }}
                </p>
                <button class="text-purple-900" @click="showLess = false">Show more</button>
              </div>
              <div v-else>
                <p>
                    {{ body }}
                </p>
                <button @click="showLess = true">Show less</button>
              </div>
            </div>
          </div>
      </div>
    </div>
  </article>
</template>

<script>
export default {
    name: 'Result',
    data() {
        return {
          showLess: true
        }
    },
    props: {
        from: String,
        to: String,
        date: String,
        subject: String,
        body: String
    }
}
</script>

<style scoped>
.message-container {
  white-space: pre-line;
}
</style>