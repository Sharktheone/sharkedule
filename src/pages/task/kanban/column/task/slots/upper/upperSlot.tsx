import {Slot} from "@kanban/column/task/slots/slotTypes"
import {useContext} from "react"
import {SlotContext} from "@kanban/column/task/slots/slotProvider"
import styles from "@kanban/column/task/slots/styles.module.scss"
import Render from "@kanban/column/task/slots/render/render"


export default function UpperSlot() {
    const context = useContext(SlotContext)
    let upperSlot = context?.upperSlot
    if (!upperSlot) return null

    if (upperSlot.length === 0) return null


    return (
        <div className={styles.slot}>
            {upperSlot.map((slot: Slot) => (
                <Render key={slot.type} slot={slot}/>
            ))}
        </div>
    )

}