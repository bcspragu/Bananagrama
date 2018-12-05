import Vue from 'vue'
import {BananaServiceClient} from '@/proto/banana_pb_service';

declare module 'vue/types/vue' {
  interface Vue {
    $client: BananaServiceClient
  }
}
