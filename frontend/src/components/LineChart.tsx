import { Line } from "react-chartjs-2";
import { observation } from "../interfaces";

export default function LineChart(props: { observations: observation[] }) {
    const formatLastObservationInstant = (ob: observation): string => {
        let lastObservationDate = new Date(ob.instant).toLocaleDateString();
        let lastObservationTime = new Date(ob.instant).toLocaleTimeString();
        return `${lastObservationDate}, ${lastObservationTime}`;
    };

    const labels = props.observations.map(ob => formatLastObservationInstant(ob));

    const data = {
        labels,
        datasets: [
            {
                label: props.observations.at(0)?.trackerName,
                data: props.observations.map(ob => ob.value),
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            },
        ],
    };

    const options = {
        responsive: true,
        plugins: {
            legend: {
                position: 'top' as const,
            },
            title: {
                display: true,
                text: `${props.observations.length} observations over time`,
            },
        },
        scales: {
            x: {
                ticks: {
                    display: false
                }
            }
        }
    };

    return (
        <Line
            options={options}
            data={data}
            style={{ maxHeight: "50vh", margin: "0 auto", marginBottom: "50px", width: "90vw" }}
        />
    )
}