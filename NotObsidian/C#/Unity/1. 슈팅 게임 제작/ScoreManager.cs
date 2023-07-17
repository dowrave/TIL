using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class ScoreManager : MonoBehaviour
{

    public Text currentScoreUI;
    private int CurrentScore;
    public Text bestScoreUI;
    private int bestScore;

    public void SetScore(int value)
    {
        // 3. 값을 할당한다
        CurrentScore = value;

        // 4. 점수 표시하기
        currentScoreUI.text = "현재 점수 : " + CurrentScore;

        if (CurrentScore > bestScore)
        {
            bestScore = CurrentScore;
            bestScoreUI.text = "최고 점수 : " + bestScore;

            // 점수 저장하기
            PlayerPrefs.SetInt("Best Score", bestScore);
        }

    }

    public int GetScore()
    {
        return CurrentScore;
    }
    void Start()
    {
        // 1. 최고 점수 불러오기
        bestScore = PlayerPrefs.GetInt("Best Score", 0);
        // 2. 최고 점수 표시하기
        bestScoreUI.text = "최고 점수 : " + bestScore;
    }
}
