package com.github.mdeluise.tracky.stats.tracker;

import com.github.mdeluise.tracky.observation.Observation;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class Processor {
    public Number getMean(List<Observation> observations) {
        float result = 0;
        for (Observation observation : observations) {
            result += observation.getValue().floatValue();
        }
        return result / observations.size();
    }


    public Number getMeanAtLastValues(int lastValues, List<Observation> observations) {
        return getMean(getLastElements(lastValues, observations));
    }


    public Number getMax(List<Observation> observations) {
        return observations.stream()
                           .max((ob1, ob2) -> Float.compare(ob1.getValue().floatValue(), ob2.getValue().floatValue()))
                           .get().getValue();
    }


    public Number getMax(int lastValues, List<Observation> observations) {
        return getMax(getLastElements(lastValues, observations));
    }


    public Number getMin(List<Observation> observations) {
        return observations.stream()
                           .min((ob1, ob2) -> Float.compare(ob1.getValue().floatValue(), ob2.getValue().floatValue()))
                           .get().getValue();
    }


    public Number getMin(int lastValues, List<Observation> observations) {
        return getMin(getLastElements(lastValues, observations));
    }


    private List<Observation> getLastElements(int lastValues, List<Observation> observations) {
        return observations.subList(Math.max(0, observations.size() - lastValues), observations.size());
    }
}
