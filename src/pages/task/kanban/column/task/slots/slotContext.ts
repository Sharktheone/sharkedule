import {Slot} from "@kanban/column/task/slots/slotTypes"
import {createContext} from "react"

export type SlotContextType = {
    upperSlot: Slot[] | null
    lowerSlot: Slot[] | null
    border: string | null
    color: string | null
}

export const SlotContext = createContext<SlotContextType | undefined>(undefined)