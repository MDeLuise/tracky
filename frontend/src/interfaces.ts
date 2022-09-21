export interface tracker {
    id: number,
    name: string,
    description: string,
    unit: string,
    lastObservation?: observation,
}

export interface observation {
    id: number,
    trackerId: number,
    trackerName: string,
    unit: string,
    instant: Date,
    value: number,
    note?: string,
}