import type {Note} from "$lib/models/note";

export interface PersonQuality {
  uuid: string
  name: string
  type: string
  createdAt: string
  notes?: Note[]
}