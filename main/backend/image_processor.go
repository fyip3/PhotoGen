package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "sync"
)

func generateImages(description string) ([]GeneratedImage, error) {
    predefinedStyles := []string{
        "professional headshot",
        "modern style",
        "formal attire",
        "warm lighting",
        "outdoor background",
    }

    var wg sync.WaitGroup
    var mu sync.Mutex
    results := []GeneratedImage{}

    for _, style := range predefinedStyles {
        wg.Add(1)
        go func(style string) {
            defer wg.Done()
            // Generate an image using the description and the current style
            generatedImage, err := callAIAPI(description)
            if err != nil {
                return 
            }

            mu.Lock()
            results = append(results, generatedImage)
            mu.Unlock()
        }(style)
    }

    wg.Wait()
    return results, nil
}


func callAIAPI(prompt string) (GeneratedImage, error) {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        return GeneratedImage{}, fmt.Errorf("OpenAI API key not set")
    }

    url := "https://api.openai.com/v1/images/generations"
    requestBody, _ := json.Marshal(map[string]interface{}{
        "prompt": prompt,
        "n":      1,
        "size":   "1024x1024",
    })

    fmt.Println("Request Body:", string(requestBody))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        return GeneratedImage{}, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return GeneratedImage{}, fmt.Errorf("failed to call OpenAI API: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("API Error Response:", string(body))  
        return GeneratedImage{}, fmt.Errorf("OpenAI API error: %s", body)
    }

    var response struct {
        Data []struct {
            URL string `json:"url"`
        } `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return GeneratedImage{}, fmt.Errorf("failed to decode OpenAI response: %w", err)
    }

    fmt.Println("Response Data:", response)  

    if len(response.Data) == 0 {
        return GeneratedImage{}, fmt.Errorf("no images generated")
    }

    return GeneratedImage{
        URL:   response.Data[0].URL,
        Style: prompt,
    }, nil
}
