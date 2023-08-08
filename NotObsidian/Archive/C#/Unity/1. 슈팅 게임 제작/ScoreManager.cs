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

    // 싱글턴 객체
    public static ScoreManager Instance = null;

    public int Score
    {
        get
        {
            return CurrentScore;
        }
        set
        {
            // 3. 값을 할당한다
            CurrentScore = value;

            // 4. 점수 표시하기
            currentScoreUI.text = "현재점수  " + CurrentScore;

            if (CurrentScore > bestScore)
            {
                bestScore = CurrentScore;
                bestScoreUI.text = "최고점수  " + bestScore;

                // 점수 저장하기
                PlayerPrefs.SetInt("Best Score", bestScore);
            }
        }
    }

    // 싱글턴 객체에 값이 없으면 생성된 자기 자신을 할당함
    void Awake()
    {
        if(Instance == null)
        {
            Instance = this;
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
