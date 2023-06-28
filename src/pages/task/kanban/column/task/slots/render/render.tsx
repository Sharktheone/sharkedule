import {Slot, SlotTypes} from "@kanban/column/task/slots/slotTypes"
import RenderTags from "@kanban/column/task/slots/render/tags/renderTag"


type Props = {
    slot: Slot
}

export default function Render({slot}: Props) {
    switch (slot.type) {
        case SlotTypes.TAGS:
            if (!slot.value) return null
            return <RenderTags tags={slot.value as string[]}/>
        default:
            return null
    }
}