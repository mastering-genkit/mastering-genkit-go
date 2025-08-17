flowchart LR
    subgraph Input["Multimodal Inputs"]
        IMG[Image]
        AUD[Audio]
        TXT[Text]
    end
    
    subgraph Process["Genkit Processing"]
        COMBINE[Combine<br/>Inputs]
        SCHEMA[Apply<br/>Schema]
        REQUEST[Model<br/>Request]
    end
    
    subgraph Output["Structured Output"]
        JSON[JSON<br/>Response]
        TYPED[Type-Safe<br/>Data]
    end
    
    IMG --> COMBINE
    AUD --> COMBINE
    TXT --> COMBINE
    COMBINE --> SCHEMA
    SCHEMA --> REQUEST
    REQUEST --> JSON
    JSON --> TYPED