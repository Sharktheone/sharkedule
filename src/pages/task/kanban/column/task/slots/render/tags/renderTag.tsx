import {TagsSlot} from "@kanban/column/task/slots/slotTypes"
import {kanbanTagType} from "@kanban/types"
import styles from "@kanban/column/task/slots/render/tags/styles.module.scss"
import {getTag} from "@/pages/task/utils/tag"
import {Tag} from "@kanban/types2"


type Props = {
    tagSlot: TagsSlot
}

export default function RenderTags({tagSlot}: Props) {
    let tagsUUIDs = tagSlot as string[]

    let tags = tagsUUIDs.map((uuid) => {
        return getTag(uuid) as Tag
    })



    return (
        <div>
            {tags?.map((tag: Tag) => (
                <RenderTag key={tag.uuid} tag={tag}/>
            ))}
        </div>
    )
}

type RenderTagProps = {
    tag: kanbanTagType
}

export function RenderTag({tag}: RenderTagProps) {
    return (
        <div className={styles.tag} style={{backgroundColor: `${tag.color}80`, borderColor: tag.color}}>
            {tag.name}
        </div>
    )
}