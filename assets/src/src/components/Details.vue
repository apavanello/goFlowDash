<script lang="ts">



export default {
  data () {
    return {
      btn: null
    }
  },
  methods: {
    openDialog()  {


      console.log('openDialog')

    }
  }
}


</script>

<template>
  <v-dialog

      width="800"
  >
    <v-card
        prepend-icon="mdi-account"
        title="User Profile"
    >
      <v-card-text>
        <v-row dense>
          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Node Label*"
                required
            ></v-text-field>
          </v-col>
          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Descrição"
            ></v-text-field>
          </v-col>

          <v-col
              cols="12"
              sm="12"
          >
            <v-text-field
                label="Repositorio"
            ></v-text-field>
          </v-col>
        </v-row>

        <small class="text-caption text-medium-emphasis">*indicates required field</small>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn
            text="Close"
            variant="plain"

        ></v-btn>

        <v-btn
            color="primary"
            text="Save"
            variant="tonal"

        ></v-btn>
      </v-card-actions>

      </v-card>
  </v-dialog>
</template>

<style scoped>

</style>
<style scoped>

</style>