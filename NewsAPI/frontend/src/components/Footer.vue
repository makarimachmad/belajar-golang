<template>
  <v-card height="400px">
    <v-footer
      v-bind="localAttrs"
      :padless="padless"
    >
      <v-card
        flat
        tile
        width="100%"
        class="green lighten-1 text-center"
      >
        <v-card-text>
          <v-btn
            v-for="icon in icons"
            :key="icon"
            class="mx-4"
            icon
          >
            <v-icon size="24px">
              {{ icon }}
            </v-icon>
          </v-btn>
        </v-card-text>

        <!-- <v-divider></v-divider> -->

        <v-card-text class="white--text">
          <h3>SPE News</h3>{{ new Date().getFullYear() }} — <strong>makarimachmad</strong>
        </v-card-text>
      </v-card>
    </v-footer>

    <!-- <v-row
      align="center"
      justify="center"
      class="ma-12"
    >
      <v-col
        cols="12"
        md="8"
      >
        <v-select
          v-model="variant"
          :items="items"
          clearable
          label="Variant"
        ></v-select>

        <v-checkbox
          v-model="padless"
          hide-details
          label="Padless"
        ></v-checkbox>
      </v-col>
    </v-row> -->
  </v-card>
</template>
<script>
  export default {
    data: () => ({
      icons: [
        'mdi-instagram',
        'mdi-email',
        'mdi-facebook',
        'mdi-twitter',
      ],
      items: [
        'default',
        'absolute',
        'fixed',
      ],
      padless: true,
      variant: 'absolute',
    }),
    computed: {
      localAttrs () {
        const attrs = {}

        if (this.variant === 'default') {
          attrs.absolute = false
          attrs.fixed = false
        } else {
          attrs[this.variant] = true
        }
        return attrs
      },
    },
  }
</script>