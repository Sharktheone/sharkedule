import {createContext, Dispatch, ReactNode, SetStateAction, useContext, useMemo, useState} from "react"
import {
    Configuration,
    DateDueSlot,
    IndexedSlot,
    PrioritySlot,
    ProgressSlot,
    Slot,
    SlotColors,
    SlotNames,
    StageSlot,
    StatusSlot,
} from "@kanban/column/task/slots/slotTypes"
import {Task} from "@kanban/types2"
import {EnvironmentContext} from "@kanban/environment"
import {getTask} from "@/pages/task/utils/task"

type SlotContextType = {
    upperSlot: Slot[] | null
    lowerSlot: Slot[] | null
    border: string | null
    color: string | null
}

export const SlotContext = createContext<SlotContextType | undefined>(undefined)

type Props = {
    children: ReactNode
    task: string
}

const config: Configuration = {
    upper: [
        SlotNames.TAGS,
        SlotNames.PRIORITY,
        SlotNames.STATUS,
        SlotNames.STAGE,
        SlotNames.PROGRESS,
    ],
    lower: [
        SlotNames.DATE_DUE,
        SlotNames.MEMBERS,
        SlotNames.SUBTASKS,
        SlotNames.CUSTOM_FIELDS,
        SlotNames.CHECKLIST,
    ]
}


// TODO: This method of rendering tags etc is not very efficient, as it requires a lot of looping over the same data.
//  I'm a lazy b... , so I'll leave it for now, but maybe in the year 3048 or something I'll fix it - or may not KEKW.
export function SlotProvider({children, task}: Props) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)

    const [t, setT] = useState<Task | undefined>(() => getTask(task))

    function slotify() {
        let upperSlot: Slot[] = []
        let lowerSlot: Slot[] = []
        let border: string | null = null
        let color: string | null = null
        let slots: IndexedSlot = {} as IndexedSlot

        if (t?.tags) {
            slots.tags = t?.tags
        }
        if (t?.priority) {
            slots.priority = t?.priority
        }
        if (t?.status) {
            slots.status = t?.status
        }
        if (t?.due_date) {
            slots.date_due = t?.due_date
        }
        if (t?.stage) {
            slots.stage = t?.stage
        }
        if (t?.members) {
            slots.members = t?.members
        }
        if (t?.progress) {
            slots.progress = t?.progress
        }
        if (t?.subtasks) {
            slots.subtasks = t?.subtasks
        }
        if (t?.custom_fields) {
            slots.custom_fields = t?.custom_fields
        }
        if (t?.checklists) {
            slots.checklists = t?.checklists
        }

        for (let slot in config.lower) {
            const slotName = config.lower[slot]
            const s: Slot = slots[slotName]
            if (s) lowerSlot.push(s)
        }

        for (let slot in config.upper) {
            const slotName = config.upper[slot]
            const s: Slot = slots[slotName]
            if (s) upperSlot.push(s)
        }

        if (config.border) {
            const b = config.border
            const s = slots[b]
            border = getSlotColor(s)
        }

        if (config.color) {
            const c = config.color
            const s = slots[c]
            color = getSlotColor(s)
        }

        return {upperSlot, lowerSlot, border, color}
    }

    function getSlotColor(slot: Slot): string | null {
        switch (slot.type) {
            case SlotColors.PRIORITY: {
                const prioritySlot = slot as PrioritySlot
                if (prioritySlot.priority.color) return prioritySlot.priority.color
                return null
            }
            case SlotColors.STATUS: {
                const statusSlot = slot as StatusSlot
                if (statusSlot.status.color) return statusSlot.status.color
                return null
            }
            case SlotColors.DATE_DUE: {
                const dateDueSlot = slot as DateDueSlot
                const dateDue = dateDueSlot.due_date.date
                const date = new Date(dateDue)
                const today = new Date()
                if (date > today) return "#00ff00" // TODO: Make configurable

                if (date < today) return "#ff0000"
                return null
            }

            case SlotColors.STAGE: {
                const stageSlot = slot as StageSlot
                if (stageSlot.stage.color) return stageSlot.stage.color
                return null
            }

            case SlotColors.PROGRESS: {
                const progressSlot = slot as ProgressSlot
                const progress = progressSlot.progress.percentage
                if (progress >= 100) return "#00ff00" // TODO: Make configurable
                if (progress > 80) return "#d5ff18"
                if (progress > 70) return "#ffff00"
                if (progress > 60) return "#ffcc00"
                if (progress > 30) return "#ff9900"
                if (progress > 20) return "#ff6600"
                if (progress > 10) return "#ff3300"
                return "#ff0000"
            }

            case SlotColors.CUSTOM_FIELDS: {
                return "#4433ff" // TODO: Make configurable which field to use
            }

            case SlotColors.CHECKLIST: {
                return "#4433ff" // TODO: Make configurable which checklist to use
            }

            default:
                return null
        }
    }

    return (
        <SlotContext.Provider value={slotify()}>
            {children}
        </SlotContext.Provider>
    )
}